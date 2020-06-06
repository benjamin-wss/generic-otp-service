package utilities

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type EnvironmentVariableUtilities struct{}

func (instance EnvironmentVariableUtilities) GetEnvironmentVariableValuePointer(fieldName string) *string {
	value := os.Getenv(fieldName)

	if len(strings.TrimSpace(value)) < 1 {
		return nil
	}

	trimmedValue := strings.TrimSpace(value)

	return &trimmedValue
}

func (instance EnvironmentVariableUtilities) GetEnvironmentVariableAsString(fieldName string) string {
	value := instance.GetEnvironmentVariableValuePointer(fieldName)

	if value == nil {
		panicMessage := fmt.Sprintf("%s environment variable not specified.", fieldName)
		panic(panicMessage)
	}

	return *value
}

func (instance EnvironmentVariableUtilities) GetEnvironmentVariableAsInteger(fieldName string) int {
	stringValue := instance.GetEnvironmentVariableAsString(fieldName)

	value, intCastingError := strconv.Atoi(stringValue)

	if intCastingError != nil {
		panicMessage := fmt.Sprintf("Error parsing %s environment variable as int.", fieldName)
		panic(panicMessage)
	}

	return value
}

func (instance EnvironmentVariableUtilities) GetEnvironmentVariableAsBoolean(fieldName string) bool {
	integerValue := instance.GetEnvironmentVariableAsInteger(fieldName)

	if integerValue == 1 {
		return true
	}

	if integerValue == 0 {
		return false
	}

	panicMessage := fmt.Sprintf("Error parsing %s environment variable as boolean. Value must be 1 for true or 0 for false.", fieldName)
	panic(panicMessage)
}
