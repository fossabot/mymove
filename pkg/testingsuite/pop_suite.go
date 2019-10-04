package testingsuite

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	envy "github.com/codegangsta/envy/lib"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gofrs/flock"
	"github.com/pkg/errors"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"0123456789"

var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

var fileLock = flock.New(os.TempDir() + "/server-test-lock.lock")

// StringWithCharset returns a random string
// https://www.calhoun.io/creating-random-strings-in-go/
func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// PopTestSuite is a suite for testing
type PopTestSuite struct {
	BaseTestSuite
	PackageName
	db     *pop.Connection
	dbName string
}

func commandWithDefaults(command string, args ...string) *exec.Cmd {
	host := envy.MustGet("DB_HOST")
	port := envy.MustGet("DB_PORT_TEST")
	user := envy.MustGet("DB_USER")
	// Get password to ensure it is set in the environment, otherwise commands won't work
	envy.MustGet("DB_PASSWORD")

	defaults := []string{"-U", user, "-h", host, "-p", port}

	arguments := append(defaults, args...)

	// #nosec G204
	return exec.Command(command, arguments...)
}

func runCommand(cmd *exec.Cmd, desc string) ([]byte, error) {
	cmdErr := bytes.Buffer{}
	cmd.Stderr = &cmdErr
	out, err := cmd.Output()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to %s: ran %s; got %s", desc, string(out), cmdErr.String())
	}
	return out, nil
}

func dropDB(destination string) error {
	drop := commandWithDefaults("dropdb", "--if-exists", destination)
	if _, err := runCommand(drop, "drop the database"); err != nil {
		return err
	}
	return nil
}

func cloneDatabase(source, destination string) error {
	lockCtx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()

	_, lockErr := fileLock.TryLockContext(lockCtx, 678*time.Millisecond)
	if lockErr != nil {
		return lockErr
	}

	if err := dropDB(destination); err != nil {
		return err
	}

	create := commandWithDefaults("createdb", destination)
	if _, err := runCommand(create, "create the database"); err != nil {
		return err
	}

	dump := commandWithDefaults("pg_dump", source)
	out, dumpErr := runCommand(dump, "dump the database")
	if dumpErr != nil {
		return dumpErr
	}

	restore := commandWithDefaults("psql", "-q", destination)
	restore.Stdin = bytes.NewReader(out)

	if _, err := runCommand(restore, "import the dump with psql"); err != nil {
		return dumpErr
	}

	if err := fileLock.Unlock(); err != nil {
		return err
	}

	return nil
}

// PackageName represents the project-relative name of a Go package.
type PackageName string

func (pn PackageName) String() string {
	return string(pn)
}

// Suffix returns a new PackageName with an underscore and the suffix appended to the end
// suffix should be a snake case string
func (pn PackageName) Suffix(suffix string) PackageName {
	return PackageName(pn.String() + "_" + suffix)
}

// CurrentPackage returns the project-relative name of the caller's package.
//
// "github.com/transcom/mymove/pkg/" is removed from the beginning of the absolute package name, so
// the return value will be e.g. "handlers/internalapi".
func CurrentPackage() PackageName {
	pc, _, _, _ := runtime.Caller(1)
	caller := runtime.FuncForPC(pc)

	fnName := strings.Replace(caller.Name(), "github.com/transcom/mymove/pkg/", "", 1)
	pkg := strings.Split(fnName, ".")[0]
	return PackageName(pkg)
}

// NewPopTestSuite returns a new PopTestSuite
func NewPopTestSuite(packageName PackageName) PopTestSuite {
	uniq := StringWithCharset(6, charset)
	dbName := fmt.Sprintf("test_%s_%s", strings.Replace(packageName.String(), "/", "_", -1), uniq)
	log.Printf("package %s is attempting to connect to database %s", packageName.String(), dbName)

	fmt.Printf("attempting to clone database %s to %s... ", "test_db", dbName)
	if err := cloneDatabase("test_db", dbName); err != nil {
		log.Panicf("failed to clone database '%s' to '%s': %#v", "testdb", dbName, err)
	}
	fmt.Println("success")

	conn, err := pop.NewConnection(&pop.ConnectionDetails{
		Dialect:  "postgres",
		Database: dbName,
		Host:     envy.MustGet("DB_HOST"),
		Port:     envy.MustGet("DB_PORT_TEST"),
		User:     envy.MustGet("DB_USER"),
		Password: envy.MustGet("DB_PASSWORD"),
	})
	if err != nil {
		log.Panic(err)
	}

	if err := conn.Open(); err != nil {
		log.Panic(err)
	}

	return PopTestSuite{db: conn, dbName: dbName, PackageName: packageName}
}

// DB returns a db connection
func (suite *PopTestSuite) DB() *pop.Connection {
	return suite.db
}

// MustSave requires saving without errors
func (suite *PopTestSuite) MustSave(model interface{}) {
	t := suite.T()
	t.Helper()

	verrs, err := suite.db.ValidateAndSave(model)
	if err != nil {
		suite.T().Errorf("Errors encountered saving %v: %v", model, err)
	}
	if verrs.HasAny() {
		suite.T().Errorf("Validation errors encountered saving %v: %v", model, verrs)
	}
}

// MustCreate requires creating without errors
func (suite *PopTestSuite) MustCreate(db *pop.Connection, model interface{}) {
	t := suite.T()
	t.Helper()

	verrs, err := db.ValidateAndCreate(model)
	if err != nil {
		suite.T().Errorf("Errors encountered creating %v: %v", model, err)
	}
	if verrs.HasAny() {
		suite.T().Errorf("Validation errors encountered creating %v: %v", model, verrs)
	}
}

// MustDestroy requires deleting without errors
func (suite *PopTestSuite) MustDestroy(model interface{}) {
	t := suite.T()
	t.Helper()

	err := suite.db.Destroy(model)
	if err != nil {
		suite.T().Errorf("Errors encountered destroying %v: %v", model, err)
	}
}

// NoVerrs prints any errors it receives
func (suite *PopTestSuite) NoVerrs(verrs *validate.Errors) bool {
	if !suite.False(verrs.HasAny()) {
		fmt.Println(verrs.String())
		return false
	}
	return true
}

// TearDown runs the teardown for step for the suite
func (suite *PopTestSuite) TearDown() {
	// disconnect other users
	if err := suite.db.Close(); err != nil {
		log.Panic(err)
	}
	// Remove the test DB
	if err := dropDB(suite.dbName); err != nil {
		log.Panic(err)
	}
}
