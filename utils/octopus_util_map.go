package utils

import (
	"reflect"
	"strings"
)

func CompareMap(defaultMap map[interface{}]interface{}, newMap map[interface{}]interface{}) map[interface{}]interface{} {
	if newMap == nil || len(newMap) == 0 {
		return defaultMap
	}
	if defaultMap == nil || len(defaultMap) == 0 {
		return newMap
	}
	for key, val := range newMap {
		if val == nil {
			continue
		}
		value := val
		queryvalue := defaultMap[key]
		if queryvalue != nil && queryvalue != "" {
			valueType := reflect.TypeOf(value)
			if strings.Contains(valueType.String(), "map") {
				valuemap := value.(map[interface{}]interface{})
				querymap := queryvalue.(map[interface{}]interface{})
				CompareMap(querymap, valuemap)
			} else if value != queryvalue {
				(defaultMap)[key] = val
			}
		} else {
			(defaultMap)[key] = val
		}
	}
	return defaultMap
}
