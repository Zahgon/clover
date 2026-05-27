package internal

import (
	"reflect"
)

type Value struct {
	V interface{}
}

func processStructTag(tagStr string) (string, bool) { _ = "STUB: not implemented"; return "", false }

// when tagStr is "", tags[0] will also be ""

func isEmptyValue(v reflect.Value) bool { _ = "STUB: not implemented"; return false }

func normalizeStruct(structValue reflect.Value) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func normalizeSlice(sliceValue reflect.Value) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getElemValueAndType(v interface{}) (reflect.Value, reflect.Type) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), *new(reflect.Type)
}

func normalizeMap(mapValue reflect.Value) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Normalize(value interface{}) (interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

func createRenameMap(rv reflect.Value, renameMap map[string]string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func rename(fields map[string]interface{}, v interface{}) map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}

func getElemType(rt reflect.Type) reflect.Type {
	_ = "STUB: not implemented"
	return *new(reflect.Type)
}

func renameMapKeys(m map[string]interface{}, v interface{}) map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}

func Encode(v map[string]interface{}) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func Decode(data []byte, m *map[string]interface{}) error { _ = "STUB: not implemented"; return nil }

func Convert(m map[string]interface{}, v interface{}) error { _ = "STUB: not implemented"; return nil }
