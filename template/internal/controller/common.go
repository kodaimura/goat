package controller

import (
	"reflect"
    "regexp"
    "strings"
	"encoding/json"

	"goat/internal/core/errs"
)

func NewBindError(err error, dataStruct interface{}) error {
    if err == nil {
        return nil
    }

	if jsonErr, ok := err.(*json.UnmarshalTypeError); ok {
		return errs.NewBadRequestError(jsonErr.Field)
	}
	if _, ok := err.(*json.SyntaxError); ok {
		return errs.NewBadRequestError("")
	}
	if strings.Contains(err.Error(), "Key:") {
		fieldName := extractFieldName(err.Error())
		return errs.NewBadRequestError(getFieldJsonTag(dataStruct, fieldName))
	}

    return errs.NewBadRequestError("")
}


func getFieldJsonTag(dataStruct interface{}, fieldName string) string {
    val := reflect.TypeOf(dataStruct).Elem()

    for i := 0; i < val.NumField(); i++ {
        field := val.Field(i)
        if field.Name == fieldName {
			jsonTag := field.Tag.Get("json")
			if jsonTag != "" && jsonTag != "-" {
				return jsonTag
			}
			return fieldName
		}
    }
    return fieldName
}

func extractFieldName(errorMsg string) string {
	re := regexp.MustCompile(`Key:\s*'([^']+)'`)
	match := re.FindStringSubmatch(errorMsg)
	if len(match) > 1 {
		return strings.Split(match[1], ".")[1]
	}
	return ""
}