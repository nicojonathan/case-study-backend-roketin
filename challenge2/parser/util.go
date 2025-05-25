package parser

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func formParser(r *http.Request, v interface{}) error {
	err := r.ParseForm()
	if err != nil {
		fmt.Println("e:parseForm", err)
		return err
	}

	var reflectValue = reflect.ValueOf(v)

	if reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}

	var reflectType = reflectValue.Type()

	forms := r.Form
	mapEntity := make(map[string]interface{})

	for formName := range forms {
		for i := 0; i < reflectValue.NumField(); i++ {
			fieldName := reflectType.Field(i).Tag.Get("json")
			fieldType := reflectType.Field(i).Type.String()
			fieldName = strings.ReplaceAll(fieldName, ",omitempty", "")
			dataType := reflectType.Field(i).Tag.Get("type")

			if fieldName == formName {
				value := forms[formName][0]
				if fieldType == "int" || dataType == "int" {
					mapEntity[formName], _ = strconv.Atoi(value)
				} else if fieldType == "float32" || dataType == "float32" {
					mapEntity[formName], _ = strconv.ParseFloat(value, 32)
				} else if fieldType == "float64" || dataType == "float64" {
					mapEntity[formName], _ = strconv.ParseFloat(value, 64)
				} else if fieldType == "time.Duration" || dataType == "time.Duration" {
					duration, _ := time.ParseDuration(value)
					mapEntity[formName] = duration
				} else if fieldType == "bool" || dataType == "bool" {
					mapEntity[formName] = value == "1" || strings.ToLower(value) == "true"
				} else {
					mapEntity[formName] = value
				}
			}
		}
	}

	jsonEntity, err := json.Marshal(mapEntity)
	if err != nil {
		fmt.Println("e:jsonMarshal", err)
		return err
	}

	err = json.Unmarshal(jsonEntity, &v)
	if err != nil {
		fmt.Println("e:jsonUnmarshal", err)
		return err
	}

	return nil
}

func ParseIDs(idsStr string) ([]int, error) {
	idStrings := strings.Split(idsStr, ",")
	var ids []int

	for _, s := range idStrings {
		s = strings.TrimSpace(s)
		if s == "" {
			continue
		}
		id, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}

	return ids, nil
}

func ParsePaginationParams(r *http.Request) (page int, limit int, err error) {
	query := r.URL.Query()

	// Default values if not provided
	page = 1
	limit = 10

	if pageStr := query.Get("page"); pageStr != "" {
		page, err = strconv.Atoi(pageStr)
		if err != nil || page < 1 {
			return 0, 0, fmt.Errorf("invalid page number")
		}
	}

	if limitStr := query.Get("limit"); limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil || limit < 1 {
			return 0, 0, fmt.Errorf("invalid limit")
		}
	}

	return page, limit, nil
}
