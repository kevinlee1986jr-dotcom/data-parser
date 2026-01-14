package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
)

// FileExists checks if a file exists
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// ReadFile reads the contents of a file
func ReadFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

// WriteFile writes data to a file
func WriteFile(filename string, data []byte) error {
	return ioutil.WriteFile(filename, data, 0644)
}

// ParseJSON parses a JSON string
func ParseJSON(data string, target interface{}) error {
	return json.Unmarshal([]byte(data), target)
}

// ToInt converts a string to an integer
func ToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

// ToFloat converts a string to a float64
func ToFloat(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

// GetEnv gets the value of an environment variable
func GetEnv(key string) string {
	return os.Getenv(key)
}

// SetEnv sets the value of an environment variable
func SetEnv(key string, value string) error {
	return os.Setenv(key, value)
}

// GetAbsolutePath gets the absolute path of a file or directory
func GetAbsolutePath(path string) (string, error) {
	return filepath.Abs(path)
}

// IsEmpty checks if a string, slice, map, or array is empty
func IsEmpty(v interface{}) bool {
	if v == nil {
		return true
	}
	switch reflect.TypeOf(v).Kind() {
	case reflect.String, reflect.Slice, reflect.Map, reflect.Array:
		return reflect.ValueOf(v).Len() == 0
	default:
		return false
	}
}

// SplitAndTrim splits a string and trims each part
func SplitAndTrim(s string, sep string) []string {
	parts := strings.Split(s, sep)
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}
	return parts
}

// ErrInvalidType is an error for invalid type
var ErrInvalidType = errors.New("invalid type")

// ErrFileNotFound is an error for file not found
var ErrFileNotFound = errors.New("file not found")

// ErrInvalidJSON is an error for invalid JSON
var ErrInvalidJSON = errors.New("invalid JSON")