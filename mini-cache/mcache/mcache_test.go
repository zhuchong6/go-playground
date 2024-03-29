package mcache

import (
	"reflect"
	"testing"
)

func TestGetterFunc_Get(t *testing.T) {
	var f GetterFunc = GetterFunc(func(key string) ([]byte, error) {
		return []byte(key), nil
	})

	expect := []byte("key")
	if v, _ := f.Get("key"); !reflect.DeepEqual(v, expect) {
		t.Errorf("callback failed")
	}
}
