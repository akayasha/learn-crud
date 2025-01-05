package utils

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
)

type StandardResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func Respond(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, StandardResponse{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

func RespondError(c *gin.Context, code int, message string) {
	c.JSON(code, StandardResponse{
		Code:    code,
		Message: message,
		Data:    nil,
	})
}

// ValidateStruct checks if any field in the struct is null or empty
func ValidateStruct(input interface{}) string {
	var missingFields []string
	val := reflect.ValueOf(input)
	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		value := val.Field(i)

		// Check if the field is a string and is empty
		if value.Kind() == reflect.String && strings.TrimSpace(value.String()) == "" {
			missingFields = append(missingFields, field.Name)
		}
	}

	if len(missingFields) > 0 {
		return fmt.Sprintf("Missing or empty fields: %v", strings.Join(missingFields, ", "))
	}
	return ""
}
