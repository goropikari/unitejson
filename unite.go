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
			av := a[k].([]any)
			av = append(av, x.Interface().([]any)...)
			a[k] = av
		case reflect.Map:
			av := a[k].(map[string]any)
			merge(av, x.Interface().(map[string]any))
		default:
			a[k] = v
		}
	}
}

func UniteJson(files [][]byte) map[string]any {
	jsons := make([]map[string]any, 0)
	for _, data := range files {
		if len(data) == 0 {
			continue
		}
		var jsonfile map[string]any
		json.Unmarshal(data, &jsonfile)
		jsons = append(jsons, jsonfile)
	}
	unitefile := map[string]any{}
	for _, js := range jsons {
		merge(unitefile, js)
	}
	return unitefile
}
