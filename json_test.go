package wire

import (
	"reflect"
	"testing"
)

func TestJSONSerialize(t *testing.T) {
	tests := []struct {
		name    string
		input   any
		want    string
		wantErr bool
	}{
		{"nil", nil, "null", false},
		{"string", "hello", `"hello"`, false},
		{"int", 42, "42", false},
		{"float", 3.14, "3.14", false},
		{"bool", true, "true", false},
		{"map", map[string]any{"a": 1}, `{"a":1}`, false},
		{"slice", []any{1, "two"}, `[1,"two"]`, false},
		{"bytes passthrough", []byte("raw data"), "raw data", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := JSON.Serialize(tt.input)
			if (err != nil) != tt.wantErr {
				t.Fatalf("Serialize() error = %v, wantErr %v", err, tt.wantErr)
			}
			if string(got) != tt.want {
				t.Errorf("Serialize() = %q, want %q", string(got), tt.want)
			}
		})
	}
}

func TestJSONSerializeString(t *testing.T) {
	tests := []struct {
		name  string
		input any
		want  string
	}{
		{"nil", nil, "null"},
		{"string", "hello", `"hello"`},
		{"int", 42, "42"},
		{"bytes passthrough", []byte("raw"), "raw"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := JSON.SerializeString(tt.input)
			if err != nil {
				t.Fatalf("SerializeString() error = %v", err)
			}
			if got != tt.want {
				t.Errorf("SerializeString() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestJSONDeserialize(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    any
		wantErr bool
	}{
		{"object", `{"a":1}`, map[string]any{"a": float64(1)}, false},
		{"array", `[1,2]`, []any{float64(1), float64(2)}, false},
		{"string", `"hello"`, "hello", false},
		{"number", `42`, float64(42), false},
		{"bool", `true`, true, false},
		{"null", `null`, nil, false},
		{"malformed", `{bad`, nil, true},
		{"not json", `hello world`, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := JSON.Deserialize([]byte(tt.input))
			if (err != nil) != tt.wantErr {
				t.Fatalf("Deserialize() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Deserialize() = %v (%T), want %v (%T)", got, got, tt.want, tt.want)
			}
		})
	}
}
