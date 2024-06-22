package value

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func AutoDetectValue(valueType, value string) (any, error) {
	switch strings.ToLower(valueType) {
	case "bool":
		switch value {
		case "true":
			return true, nil
		case "false":
			return false, nil
		default:
			return false, errors.New(fmt.Sprintf("invalid bool value: %s", value))
		}
	case "string":
		return value, nil
	case "float":
		return strconv.ParseFloat(value, 64)
	case "int":
		return strconv.ParseInt(value, 10, 64)
	case "json":
		m := make(map[string]interface{})
		err := json.Unmarshal([]byte(value), &m)
		return m, err
	default:
		return nil, errors.New("invalid toggle type")
	}
}

func ToString(valueType string, value any) (string, error) {
	switch strings.ToLower(valueType) {
	case "bool":
		return fmt.Sprintf("%v", value), nil
	case "string":
		return value.(string), nil
	case "float":
		return fmt.Sprintf("%f", value), nil
	case "int":
		return fmt.Sprintf("%d", value), nil
	case "json":
		resultJson, _ := json.Marshal(value)
		return string(resultJson), nil
	default:
		return "", errors.New("invalid toggle type")
	}
}
