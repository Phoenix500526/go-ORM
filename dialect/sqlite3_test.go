package dialect

import (
	"reflect"
	"testing"
	"time"
)

func TestDataTypeOf(t *testing.T) {
	dial := &sqlite3{}
	cases := []struct {
		Value interface{}
		Type  string
	}{
		{"Phoenix", "text"},
		{123, "integer"},
		{3.14, "real"},
		{[]int{1, 2, 3}, "blob"},
		{time.Now(), "datetime"},
	}
	for _, c := range cases {
		if typ := dial.DataTypeOf(reflect.ValueOf(c.Value)); typ != c.Type {
			t.Fatalf("expect %s, but got %s", c.Type, typ)
		}
	}
}
