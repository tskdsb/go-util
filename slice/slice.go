package util

import (
	"fmt"
	"reflect"
)

// RemoveElement remove one element by index
// Do not use RemoveElement in for-range
func RemoveElement(slice interface{}, index int) (interface{}, error) {
	value := reflect.ValueOf(slice)
	if value.Kind() != reflect.Slice {
		return nil, fmt.Errorf("not slice type")
	}

	if index < 0 {
		return nil, fmt.Errorf("index less than zero")
	}

	length := value.Len()
	if index >= length {
		return nil, fmt.Errorf("index greater than length")
	}

	appendSlice := reflect.AppendSlice(value.Slice(0, index), value.Slice(index+1, length))
	return appendSlice.Interface(), nil
}

// FilterSlice return a new slice contains elements checked by filter
// Example: [1,2,3] -> filter(>1) -> [2,3]
func FilterSlice(slice interface{}, filter func(element interface{}) bool) (interface{}, error) {
	value := reflect.ValueOf(slice)
	l := value.Len()
	results := reflect.MakeSlice(value.Type(), 0, l)

	for i := 0; i < l; i++ {
		e := value.Index(i).Interface()
		if filter(e) {
			results = reflect.Append(results, value.Index(i))
		}
	}

	return results.Interface(), nil
}
