package api

import (
	"fmt"
	"reflect"

	"github.com/gin-gonic/gin"
)

func resWithJSON(c *gin.Context, data interface{}) error {
	result := make(map[string]interface{})

	v := reflect.ValueOf(data)
	t := reflect.TypeOf(data)

	if v.Kind() != reflect.Struct {
		return fmt.Errorf("expected a struct, got %s", v.Kind())
	}

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		jsonTag := field.Tag.Get("json")
		if jsonTag != "" {
			result[jsonTag] = v.Field(i).Interface()
		}
	}

	c.JSON(200, result)

	return nil
}

func resWithJSONArray(c *gin.Context, data []interface{}) error {
	resultArray := []map[string]interface{}{}
	for _, dat := range data {
		result := make(map[string]interface{})

		v := reflect.ValueOf(dat)
		t := reflect.TypeOf(dat)

		if v.Kind() != reflect.Struct {
			return fmt.Errorf("expected a struct, got %s", v.Kind())
		}

		for i := 0; i < v.NumField(); i++ {
			field := t.Field(i)
			jsonTag := field.Tag.Get("json")
			if jsonTag != "" {
				result[jsonTag] = v.Field(i).Interface()
			}
		}

    resultArray=append(resultArray,result)
	}

	c.JSON(200, resultArray)

	return nil
}
