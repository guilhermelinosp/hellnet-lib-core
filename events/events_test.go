package events

import (
	"reflect"
	"testing"
)

// The Avro schemas in hellnet-schemas use the same field names as the JSON payloads, so every field must carry an
// avro tag equal to its json tag: the Avro codec matches struct fields by tag, not by Go name.
func TestAvroTagsMatchJSONTags(t *testing.T) {
	for _, event := range []any{OrderRequested{}, OrderAccepted{}} {
		typ := reflect.TypeOf(event)
		for i := 0; i < typ.NumField(); i++ {
			field := typ.Field(i)
			if field.Tag.Get("avro") == "" || field.Tag.Get("avro") != field.Tag.Get("json") {
				t.Errorf("%s.%s: avro tag %q, json tag %q", typ.Name(), field.Name, field.Tag.Get("avro"), field.Tag.Get("json"))
			}
		}
	}
}
