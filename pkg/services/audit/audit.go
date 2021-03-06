package audit

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/gobuffalo/flect"
	"github.com/gofrs/uuid"
	"go.uber.org/zap"

	"github.com/transcom/mymove/pkg/auth"
)

func Capture(model interface{}, payload interface{}, logger Logger, session *auth.Session, request *http.Request) ([]zap.Field, error) {
	var logItems []zap.Field
	eventType := extractEventType(request)
	msg := flect.Titleize(eventType)

	logItems = append(logItems,
		zap.String("event_type", eventType),
		zap.String("responsible_user_id", session.UserID.String()),
		zap.String("responsible_user_email", session.Email),
	)

	if session.IsAdminUser() || session.IsOfficeUser() {
		logItems = append(logItems,
			zap.String("responsible_user_name", fullName(session.FirstName, session.LastName)),
		)
	}

	t, err := validateInterface(model)
	if err == nil && reflect.ValueOf(model).IsValid() == true && reflect.ValueOf(model).IsNil() == false && reflect.ValueOf(model).IsZero() == false {
		recordType := parseRecordType(t.String())
		elem := reflect.ValueOf(model).Elem()

		var createdAt string
		if elem.FieldByName("CreatedAt").IsValid() == true {
			createdAt = elem.FieldByName("CreatedAt").Interface().(time.Time).String()
		} else {
			createdAt = time.Now().String()
		}

		var updatedAt string
		if elem.FieldByName("updatedAt").IsValid() == true {
			updatedAt = elem.FieldByName("updatedAt").Interface().(time.Time).String()
		} else {
			updatedAt = time.Now().String()
		}

		var id string
		if elem.FieldByName("ID").IsValid() == true {
			id = elem.FieldByName("ID").Interface().(uuid.UUID).String()
		} else {
			id = ""
		}

		logItems = append(logItems,
			zap.String("record_id", id),
			zap.String("record_type", recordType),
			zap.String("record_created_at", createdAt),
			zap.String("record_updated_at", updatedAt),
		)

		if payload != nil {
			_, err = validateInterface(payload)
			if err != nil {
				return nil, err
			}

			var payloadFields []string
			payloadValue := reflect.ValueOf(payload).Elem()
			for i := 0; i < payloadValue.NumField(); i++ {
				fieldFromType := payloadValue.Type().Field(i)
				fieldFromValue := payloadValue.Field(i)
				fieldName := flect.Underscore(fieldFromType.Name)

				if !fieldFromValue.IsZero() {
					payloadFields = append(payloadFields, fieldName)
				}
			}

			logItems = append(logItems, zap.String("fields_changed", strings.Join(payloadFields, ",")))

			var payloadJSON []byte
			payloadJSON, err = json.Marshal(payload)

			if err != nil {
				return nil, err
			}

			logger.Debug("Audit patch payload", zap.String("patch_payload", string(payloadJSON)))
		}
	} else {
		msg += " invalid or zero or nil model interface received from request handler"
		logItems = append(logItems,
			zap.Error(err),
		)
	}

	logger.Info(msg, logItems...)

	return logItems, nil
}

func parseRecordType(rt string) string {
	parts := strings.Split(rt, ".")

	return parts[1]
}

func fullName(first, last string) string {
	return first + " " + last
}

func validateInterface(thing interface{}) (reflect.Type, error) {
	t := reflect.TypeOf(thing)
	if t.Kind() != reflect.Ptr {
		return nil, errors.New("must pass a pointer to a struct")
	}

	t = t.Elem()
	if t.Kind() != reflect.Struct {
		return nil, errors.New("must pass a pointer to a struct")
	}

	return t, nil
}

func extractEventType(request *http.Request) string {
	path := request.URL.Path
	apiRegex := regexp.MustCompile("\\/[a-zA-Z]+\\/v1")
	uuidRegex := regexp.MustCompile("/([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}){1}") // https://adamscheller.com/regular-expressions/uuid-regex/
	cleanPath := uuidRegex.ReplaceAllString(apiRegex.ReplaceAllString(path, ""), "")
	return fmt.Sprintf("audit_%s_%s", strings.ToLower(request.Method), flect.Underscore(cleanPath))
}
