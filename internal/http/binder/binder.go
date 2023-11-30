package binder

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

// Bind binds the data from an HTTP request to a Go struct.
func Bind(r *http.Request, v interface{}) error {
	if r.Method != "POST" && r.Method != "PUT" {
		return fmt.Errorf("invalid method: %s", r.Method)
	}

	switch r.Header.Get("Content-Type") {
	case "application/json":
		return bindJSON(r, v)
	case "application/x-www-form-urlencoded":
		return bindForm(r, v)
	default:
		return fmt.Errorf("unsupported content type: %s", r.Header.Get("Content-Type"))
	}
}

// bindJSON binds JSON data from an HTTP request to a Go struct.
func bindJSON(r *http.Request, v interface{}) error {
	decoder := json.NewDecoder(r.Body)
	return decoder.Decode(v)
}

// bindForm binds form data from an HTTP request to a Go struct.
func bindForm(r *http.Request, v interface{}) error {
	values := r.Form
	return bindMap(values, v)
}

// bindMap binds a map of values to a Go struct.
func bindMap(values map[string][]string, v interface{}) error {
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("destination must be a pointer to a struct")
	}

	for key, values := range values {
		fieldName := strings.ToLower(strings.Replace(key, "-", "_", -1))
		field := val.Elem().FieldByName(fieldName)
		if !field.IsValid() {
			continue
		}

		switch field.Kind() {
		case reflect.String:
			field.SetString(values[0])
		case reflect.Int, reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8:
			value, err := strconv.ParseInt(values[0], 10, 64)
			if err != nil {
				return fmt.Errorf("failed to parse value for field %s: %w", key, err)
			}
			field.SetInt(value)
		case reflect.Float32, reflect.Float64:
			value, err := strconv.ParseFloat(values[0], 64)
			if err != nil {
				return fmt.Errorf("failed to parse value for field %s: %w", key, err)
			}
			field.SetFloat(value)
		default:
			return fmt.Errorf("unsupported field type: %s", field.Kind())
		}
	}

	return nil
}
