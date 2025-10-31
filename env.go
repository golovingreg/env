package env

import (
	"os"
	"strconv"
)

// GetString gets environment variable, or returns default value if it is not set.
func GetString(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}

	return value
}

// GetBool gets environment variable and casts is to bool.
// Returns default value if variable is not set.
// Panics if variable can't be cast to bool.
func GetBool(key string, defaultValue bool) bool {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}

	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		panic(err)
	}

	return boolValue
}

// GetInt gets environment variable and casts is to int32.
// Returns default value if variable is not set.
// Panics if variable can't be cast to int32.
func GetInt(key string, defaultValue int) int {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}

	return intValue
}

// GetInt64 gets environment variable and casts is to int64.
// Returns default value if variable is not set.
// Panics if variable can't be cast to int64.
func GetInt64(key string, defaultValue int64) int64 {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}

	int64Value, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		panic(err)
	}

	return int64Value
}

// GetFloat32 gets environment variable and casts is to float32.
// Returns default value if variable is not set.
// Panics if variable can't be cast to float32.
func GetFloat32(key string, defaultValue float32) float32 {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}

	floatValue, err := strconv.ParseFloat(value, 32)
	if err != nil {
		panic(err)
	}

	return float32(floatValue)
}

// GetFloat64 gets environment variable and casts is to float64.
// Returns default value if variable is not set.
// Panics if variable can't be cast to float64.
func GetFloat64(key string, defaultValue float64) float64 {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}

	floatValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		panic(err)
	}

	return floatValue
}
