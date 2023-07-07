package mistake

import (
	"encoding/json"
	"fmt"

	"github.com/go-playground/validator"
)

// main structure to hold M info
type M struct {
	ReturnCode int
	Details    []Detail
}

// mistake Detail
type Detail struct {
	ID      string
	Message string
}

// value collector
var valueCollector []any = []any{}

// creates a new mistake with one message
func New(statusCode int, messageID string, message string) *M {
	errorInfo := M{
		ReturnCode: statusCode,
		Details: []Detail{
			{
				ID:      messageID,
				Message: message,
			},
		},
	}

	return &errorInfo
}

// Creates a new Mistake(M) from "github.com/go-playground/validator" type
func NewStructValidation(err error, statusCode int, messageID string, message string) *M {
	eh := M{
		ReturnCode: statusCode,
		Details:    []Detail{},
	}
	if verr, ok := err.(validator.ValidationErrors); ok {
		for _, v := range verr {
			AppendValue(v.Field())
			AppendValue(v.Tag())
			newDetail := Detail{
				ID:      messageID,
				Message: Formatter(message),
			}

			eh.Details = append(eh.Details, newDetail)
		}
	}
	return &eh
}

// append value to be applied into AppendDetail method
func AppendValue(v any) {
	valueCollector = append(valueCollector, v)
}

func Formatter(message string) string {
	newMessage := fmt.Sprintf(message, valueCollector...)
	valueCollector = []any{}
	return newMessage
}

// return string error
func (e M) Error() string {
	text, _ := json.Marshal(e)
	return fmt.Sprintf("%s", text)
}
