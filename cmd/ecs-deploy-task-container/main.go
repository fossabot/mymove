package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/99designs/aws-vault/prompt"
	"github.com/99designs/aws-vault/vault"
	"github.com/99designs/keyring"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	awssession "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchevents"
	"github.com/aws/aws-sdk-go/service/ecr"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/aws/aws-sdk-go/service/rds"
	"github.com/pkg/errors"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/transcom/mymove/pkg/cli"
)

var services = []string{"app"}
var environments = []string{"prod", "staging", "experimental"}
var rules = []string{"save_fuel_price_data"}

type errInvalidAccountID struct {
	AwsAccountID string
}

func (e *errInvalidAccountID) Error() string {
	return fmt.Sprintf("invalid AWS account ID %q", e.AwsAccountID)
}

type errInvalidRegion struct {
	Region string
}

func (e *errInvalidRegion) Error() string {
	return fmt.Sprintf("invalid AWS region %q", e.Region)
}

type errInvalidService struct {
	Service string
}

func (e *errInvalidService) Error() string {
	return fmt.Sprintf("invalid AWS ECS service %q, expecting one of %q", e.Service, services)
}

type errInvalidEnvironment struct {
	Environment string
}

func (e *errInvalidEnvironment) Error() string {
	return fmt.Sprintf("invalid MilMove environment %q, expecting one of %q", e.Environment, environments)
}

type errinvalidRepositoryName struct {
	RepositoryName string
}

func (e *errinvalidRepositoryName) Error() string {
	return fmt.Sprintf("invalid AWS ECR respository name %q", e.RepositoryName)
}

type errinvalidImageTag struct {
	ImageTag string
}

func (e *errinvalidImageTag) Error() string {
	return fmt.Sprintf("invalid AWS ECR image tag %q", e.ImageTag)
}

type errInvalidRule struct {
	Rule string
}

func (e *errInvalidRule) Error() string {
	return fmt.Sprintf("invalid AWS CloudWatch Event Target rule %q", e.Rule)
}

const (
	awsAccountIDFlag         string = "aws-account-id"
	awsRegionFlag            string = "aws-region"
	awsProfileFlag           string = "aws-profile"
	awsVaultKeychainNameFlag string = "aws-vault-keychain-name"
	chamberRetriesFlag       string = "chamber-retries"
	chamberKMSKeyAliasFlag   string = "chamber-kms-key-alias"
	chamberUsePathsFlag      string = "chamber-use-paths"
	serviceFlag              string = "service"
	environmentFlag          string = "environment"
	repositoryNameFlag       string = "repository-name"
	imageTagFlag             string = "image-tag"
	ruleFlag                 string = "rule"
)

func initFlags(flag *pflag.FlagSet) {

	// AWS Vault Settings
	flag.String(awsAccountIDFlag, "", "The AWS Account ID")
	flag.String(awsRegionFlag, "us-west-2", "The AWS Region")
	flag.String(awsProfileFlag, "", "The aws-vault profile")
	flag.String(awsVaultKeychainNameFlag, "", "The aws-vault keychain name")

	// Chamber Settings
	// TODO: Add chamberBinaryFlag and set default to /bin/chamber
	flag.Int(chamberRetriesFlag, 20, "Chamber Retries")
	flag.String(chamberKMSKeyAliasFlag, "alias/aws/ssm", "Chamber KMS Key Alias")
	flag.Int(chamberUsePathsFlag, 1, "Chamber Use Paths")

	// Task Definition Settings
	flag.String(serviceFlag, "", fmt.Sprintf("The service name (choose %q)", services))
	flag.String(environmentFlag, "", fmt.Sprintf("The environment name (choose %q)", environments))
	flag.String(repositoryNameFlag, "", fmt.Sprintf("The name of the repository where the tagged image resides"))
	flag.String(imageTagFlag, "", "The name of the image tag referenced in the task definition")
	flag.String(ruleFlag, "", fmt.Sprintf("The name of the CloudWatch Event Rule targeting the Task Definition (choose %q)", rules))

	// EIA Open Data API
	// The EIA Key is set in the Local or CircleCI environment and not in Chamber.
	cli.InitEIAFlags(flag)

	// Verbose
	cli.InitVerboseFlags(flag)

	// Don't sort flags
	flag.SortFlags = false
}

func checkConfig(v *viper.Viper) error {

	awsAccountID := v.GetString(awsAccountIDFlag)
	if len(awsAccountID) == 0 {
		return errors.Wrap(&errInvalidAccountID{AwsAccountID: awsAccountID}, fmt.Sprintf("%q is invalid", awsAccountIDFlag))
	}

	regions, ok := endpoints.RegionsForService(endpoints.DefaultPartitions(), endpoints.AwsPartitionID, endpoints.EcsServiceID)
	if !ok {
		return fmt.Errorf("could not find regions for service %q", endpoints.EcsServiceID)
	}

	region := v.GetString(awsRegionFlag)
	if len(region) == 0 {
		return errors.Wrap(&errInvalidRegion{Region: region}, fmt.Sprintf("%q is invalid", awsRegionFlag))
	}

	if _, ok := regions[region]; !ok {
		return errors.Wrap(&errInvalidRegion{Region: region}, fmt.Sprintf("%q is invalid", awsRegionFlag))
	}

	chamberRetries := v.GetInt(chamberRetriesFlag)
	if chamberRetries < 1 && chamberRetries > 20 {
		return errors.New("Chamber Retries must be greater than or equal to 1 and less than or equal to 20")
	}

	chamberKMSKeyAlias := v.GetString(chamberKMSKeyAliasFlag)
	if len(chamberKMSKeyAlias) == 0 {
		return errors.New("Chamber KMS Key Alias must be set")
	}

	chamberUsePaths := v.GetInt(chamberUsePathsFlag)
	if chamberUsePaths < 1 && chamberUsePaths > 20 {
		return errors.New("Chamber Use Paths must be greater than or equal to 1 and less than or equal to 20")
	}

	serviceName := v.GetString(serviceFlag)
	if len(serviceName) == 0 {
		return errors.Wrap(&errInvalidService{Service: serviceName}, fmt.Sprintf("%q is invalid", serviceFlag))
	}
	validService := false
	for _, str := range services {
		if serviceName == str {
			validService = true
			break
		}
	}
	if !validService {
		return errors.Wrap(&errInvalidService{Service: serviceName}, fmt.Sprintf("%q is invalid", serviceFlag))
	}

	environmentName := v.GetString(environmentFlag)
	if len(environmentName) == 0 {
		return errors.Wrap(&errInvalidEnvironment{Environment: environmentName}, fmt.Sprintf("%q is invalid", environmentFlag))
	}
	validEnvironment := false
	for _, str := range environments {
		if environmentName == str {
			validEnvironment = true
			break
		}
	}
	if !validEnvironment {
		return errors.Wrap(&errInvalidEnvironment{Environment: environmentName}, fmt.Sprintf("%q is invalid", environmentFlag))
	}

	repositoryName := v.GetString(repositoryNameFlag)
	if len(repositoryName) == 0 {
		return errors.Wrap(&errinvalidRepositoryName{RepositoryName: repositoryName}, fmt.Sprintf("%q is invalid", repositoryNameFlag))
	}

	imageTag := v.GetString(imageTagFlag)
	if len(imageTag) == 0 {
		return errors.Wrap(&errinvalidImageTag{ImageTag: imageTag}, fmt.Sprintf("%q is invalid", imageTagFlag))
	}

	ruleName := v.GetString(ruleFlag)
	if len(ruleName) == 0 {
		return errors.Wrap(&errInvalidRule{Rule: ruleName}, fmt.Sprintf("%q is invalid", ruleFlag))
	}
	validRule := false
	for _, str := range rules {
		if ruleName == str {
			validRule = true
			break
		}
	}
	if !validRule {
		return errors.Wrap(&errInvalidRule{Rule: ruleName}, fmt.Sprintf("%q is invalid", ruleFlag))
	}

	err := cli.CheckEIA(v)
	if err != nil {
		return err
	}

	return nil
}

func quit(logger *log.Logger, flag *pflag.FlagSet, err error) {
	logger.Println(err.Error())
	logger.Println("Usage of ecs-service-logs:")
	if flag != nil {
		flag.PrintDefaults()
	}
	os.Exit(1)
}

// getAWSCredentials uses aws-vault to return AWS credentials
func getAWSCredentials(keychainName string, keychainProfile string) (*credentials.Credentials, error) {

	// Open the keyring which holds the credentials
	ring, _ := keyring.Open(keyring.Config{
		ServiceName:              "aws-vault",
		AllowedBackends:          []keyring.BackendType{keyring.KeychainBackend},
		KeychainName:             keychainName,
		KeychainTrustApplication: true,
	})

	// Prepare options for the vault before creating the provider
	vConfig, err := vault.LoadConfigFromEnv()
	if err != nil {
		return nil, errors.Wrap(err, "Unable to load AWS config from environment")
	}
	vOptions := vault.VaultOptions{
		Config:    vConfig,
		MfaPrompt: prompt.Method("terminal"),
	}
	vOptions = vOptions.ApplyDefaults()
	err = vOptions.Validate()
	if err != nil {
		return nil, errors.Wrap(err, "Unable to validate aws-vault options")
	}

	// Get a new provider to retrieve the credentials
	provider, err := vault.NewVaultProvider(ring, keychainProfile, vOptions)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to create aws-vault provider")
	}
	credVals, err := provider.Retrieve()
	if err != nil {
		return nil, errors.Wrap(err, "Unable to retrieve aws credentials from aws-vault")
	}
	return credentials.NewStaticCredentialsFromCreds(credVals), nil
}

func main() {
	flag := pflag.CommandLine
	initFlags(flag)
	flag.Parse(os.Args[1:])

	v := viper.New()
	v.BindPFlags(flag)
	v.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	v.AutomaticEnv()

	// Create the logger
	// Remove the prefix and any datetime data
	logger := log.New(os.Stdout, "", log.LstdFlags)

	verbose := v.GetBool(cli.VerboseFlag)
	if !verbose {
		// Disable any logging that isn't attached to the logger unless using the verbose flag
		log.SetOutput(ioutil.Discard)
		log.SetFlags(0)

		// Remove the flags for the logger
		logger.SetFlags(0)
	}

	err := checkConfig(v)
	if err != nil {
		quit(logger, flag, err)
	}

	awsRegion := v.GetString(awsRegionFlag)

	awsConfig := &aws.Config{
		Region: aws.String(awsRegion),
	}

	keychainName := v.GetString(awsVaultKeychainNameFlag)
	keychainProfile := v.GetString(awsProfileFlag)

	if len(keychainName) > 0 && len(keychainProfile) > 0 {
		creds, err := getAWSCredentials(keychainName, keychainProfile)
		if err != nil {
			quit(logger, nil, errors.Wrap(err, fmt.Sprintf("Unable to get AWS credentials from the keychain %s and profile %s", keychainName, keychainProfile)))
		}
		awsConfig.CredentialsChainVerboseErrors = aws.Bool(verbose)
		awsConfig.Credentials = creds
	}

	sess, err := awssession.NewSession(awsConfig)
	if err != nil {
		quit(logger, nil, errors.Wrap(err, "failed to create AWS session"))
	}

	// Create the Services
	serviceCloudWatchEvents := cloudwatchevents.New(sess)
	serviceECS := ecs.New(sess)
	serviceECR := ecr.New(sess)
	serviceRDS := rds.New(sess)

	// Get the current task definition (for rollback)
	ruleName := v.GetString(ruleFlag)
	targetsOutput, err := serviceCloudWatchEvents.ListTargetsByRule(&cloudwatchevents.ListTargetsByRuleInput{
		Rule: aws.String(ruleName),
	})
	if err != nil {
		quit(logger, nil, errors.Wrap(err, "error retrieving targets for rule"))
	}

	blueTarget := targetsOutput.Targets[0]
	blueTaskDefArn := *blueTarget.EcsParameters.TaskDefinitionArn
	logger.Println(fmt.Sprintf("Blue Task Def Arn: %s", blueTaskDefArn))

	// Confirm the image exists
	imageTag := v.GetString(imageTagFlag)
	registryID := v.GetString(awsAccountIDFlag)
	repositoryName := v.GetString(repositoryNameFlag)
	imageName := fmt.Sprintf("%s.dkr.ecr.%s.amazonaws.com/%s:%s", registryID, awsRegion, repositoryName, imageTag)

	_, err = serviceECR.DescribeImages(&ecr.DescribeImagesInput{
		ImageIds: []*ecr.ImageIdentifier{
			{
				ImageTag: aws.String(imageTag),
			},
		},
		RegistryId:     aws.String(registryID),
		RepositoryName: aws.String(repositoryName),
	})
	if err != nil {
		quit(logger, nil, errors.Wrapf(err, "unable retrieving image from %s", imageName))
	}

	// Get the database host using the instance identifier
	// TODO: Allow passing in from DB_HOST
	serviceName := v.GetString(serviceFlag)
	environmentName := v.GetString(environmentFlag)
	dbInstanceIdentifier := fmt.Sprintf("%s-%s", serviceName, environmentName)
	dbInstancesOutput, err := serviceRDS.DescribeDBInstances(&rds.DescribeDBInstancesInput{
		DBInstanceIdentifier: aws.String(dbInstanceIdentifier),
	})
	if err != nil {
		quit(logger, nil, errors.Wrapf(err, "error retrieving database definition for %s", dbInstanceIdentifier))
	}
	dbHost := *dbInstancesOutput.DBInstances[0].Endpoint.Address

	// Set up some key variables
	clusterName := fmt.Sprintf("%s-%s", serviceName, environmentName)
	taskRoleArn := fmt.Sprintf("ecs-task-role-%s", clusterName)
	executionRoleArn := fmt.Sprintf("ecs-task-excution-role-%s", clusterName)
	containerDefName := fmt.Sprintf("%s-tasks-%s-%s", serviceName, ruleName, environmentName)

	// familyName is the name used to register the task
	familyName := fmt.Sprintf("%s-scheduled-task-%s-%s", serviceName, ruleName, environmentName)

	// AWS Logs Group is related to the cluster and should not be changed
	awsLogsGroup := fmt.Sprintf("ecs-tasks-%s-%s", serviceName, environmentName)
	awsLogsStreamPrefix := fmt.Sprintf("%s-tasks", serviceName)

	// Chamber Settings
	chamberRetries := v.GetInt(chamberRetriesFlag)
	chamberKMSKeyAlias := v.GetString(chamberKMSKeyAliasFlag)
	chamberUsePaths := v.GetInt(chamberUsePathsFlag)

	// Tool Settings
	eiaKey := v.GetString(cli.EIAKeyFlag)
	eiaURL := v.GetString(cli.EIAURLFlag)

	// Register the new task definition
	taskDefinitionOutput, err := serviceECS.RegisterTaskDefinition(&ecs.RegisterTaskDefinitionInput{
		ContainerDefinitions: []*ecs.ContainerDefinition{
			{
				Name:      aws.String(containerDefName),
				Image:     aws.String(imageName),
				Essential: aws.Bool(true),
				EntryPoint: []*string{
					aws.String("/bin/chamber"),
					aws.String("-r"),
					aws.String(strconv.Itoa(chamberRetries)),
					aws.String("exec"),
					aws.String(clusterName),
					aws.String("--"),
					aws.String(fmt.Sprintf("/bin/%s", ruleName)),
				},
				Command: []*string{},
				Environment: []*ecs.KeyValuePair{
					&ecs.KeyValuePair{
						Name:  aws.String("ENV"),
						Value: aws.String("container"),
					},
					&ecs.KeyValuePair{
						Name:  aws.String("ENVIRONMENT"),
						Value: aws.String(environmentName),
					},
					&ecs.KeyValuePair{
						Name:  aws.String("DB_HOST"),
						Value: aws.String(dbHost),
					},
					&ecs.KeyValuePair{
						Name:  aws.String("DB_PORT"),
						Value: aws.String("5432"),
					},
					&ecs.KeyValuePair{
						Name:  aws.String("DB_USER"),
						Value: aws.String("master"),
					},
					&ecs.KeyValuePair{
						Name:  aws.String("DB_NAME"),
						Value: aws.String("app"),
					},
					&ecs.KeyValuePair{
						Name:  aws.String("DB_SSL_MODE"),
						Value: aws.String("verify-full"),
					},
					&ecs.KeyValuePair{
						Name:  aws.String("DB_SSL_ROOT_CERT"),
						Value: aws.String("/bin/rds-combined-ca-bundle.pem"),
					},
					&ecs.KeyValuePair{
						Name:  aws.String("CHAMBER_KMS_KEY_ALIAS"),
						Value: aws.String(chamberKMSKeyAlias),
					},
					&ecs.KeyValuePair{
						Name:  aws.String("CHAMBER_USE_PATHS"),
						Value: aws.String(strconv.Itoa(chamberUsePaths)),
					},
					&ecs.KeyValuePair{
						Name:  aws.String("EIA_KEY"),
						Value: aws.String(eiaKey),
					},
					&ecs.KeyValuePair{
						Name:  aws.String("EIA_URL"),
						Value: aws.String(eiaURL),
					},
				},
				LogConfiguration: &ecs.LogConfiguration{
					LogDriver: aws.String("awslogs"),
					Options: map[string]*string{
						"awslogs-group":         aws.String(awsLogsGroup),
						"awslogs-region":        aws.String(awsRegion),
						"awslogs-stream-prefix": aws.String(awsLogsStreamPrefix),
					},
				},
			},
		},
		Cpu:                     aws.String("256"),
		ExecutionRoleArn:        aws.String(executionRoleArn),
		Family:                  aws.String(familyName),
		Memory:                  aws.String("512"),
		NetworkMode:             aws.String("awsvpc"),
		RequiresCompatibilities: []*string{aws.String("FARGATE")},
		TaskRoleArn:             aws.String(taskRoleArn),
	})
	if err != nil {
		quit(logger, nil, errors.Wrap(err, "error registering new task definition"))
	}
	greenTaskDefArn := *taskDefinitionOutput.TaskDefinition.TaskDefinitionArn
	logger.Println(fmt.Sprintf("Green Task Def Arn: %s", greenTaskDefArn))

	// Update the task event target with the new task ECS parameters
	putTargetsOutput, err := serviceCloudWatchEvents.PutTargets(&cloudwatchevents.PutTargetsInput{
		Rule: aws.String(ruleName),
		Targets: []*cloudwatchevents.Target{
			{
				Id:      blueTarget.Id,
				Arn:     blueTarget.Arn,
				RoleArn: blueTarget.RoleArn,
				EcsParameters: &cloudwatchevents.EcsParameters{
					LaunchType:           aws.String("FARGATE"),
					NetworkConfiguration: blueTarget.EcsParameters.NetworkConfiguration,
					TaskCount:            aws.Int64(1),
					TaskDefinitionArn:    aws.String(greenTaskDefArn),
				},
			},
		},
	})
	if err != nil {
		quit(logger, nil, errors.Wrap(err, "Unable to put new target"))
	}
	logger.Println(putTargetsOutput)
}
