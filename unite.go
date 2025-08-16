package main

import (
	"encoding/json"
	"reflect"
)

func merge(a, b map[string]any) {
	for k, v := range b {
		if _, ok := a[k]; !ok {
			a[k] = v
			continue
		}
		x := reflect.ValueOf(v)
		switch x.Kind() {
		case reflect.Slice:
			av, ok := a[k].([]any)
			if ok {
				av = append(av, x.Interface().([]any)...)
				a[k] = av
			} else {
				a[k] = x.Interface().([]any)
			}
		case reflect.Map:
			av := a[k].(map[string]any)
			merge(av, x.Interface().(map[string]any))
		default:
			a[k] = v
		}
	}
}

func UniteJSON(files [][]byte) (map[string]any, error) {
	jsons := make([]map[string]any, 0)
	for _, data := range files {
		if len(data) == 0 {
			continue
		}
		var jsonfile map[string]any
		if err := json.Unmarshal(data, &jsonfile); err != nil {
			return nil, err
		}
		jsons = append(jsons, jsonfile)
	}
	unitefile := map[string]any{}
	for _, js := range jsons {
		merge(unitefile, js)
	}
	return unitefile, nil
}
