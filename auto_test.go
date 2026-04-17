package wire

import (
	"reflect"
	"testing"
)

func TestAutoSerialize(t *testing.T) {
	tests := []struct {
		name    string
		input   any
		want    string
		wantErr bool
	}{
		{"nil", nil, "", false},
		{"string", "hello", "hello", false},
		{"bytes", []byte("raw"), "raw", false},
		{"int", 42, "42", false},
		{"bool", true, "true", false},
		{"map", map[string]any{"a": 1}, `{"a":1}`, false},
		{"slice", []any{1, 2}, `[1,2]`, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Auto.Serialize(tt.input)
			if (err != nil) != tt.wantErr {
				t.Fatalf("Serialize() error = %v, wantErr %v", err, tt.wantErr)
			}
			if string(got) != tt.want {
				t.Errorf("Serialize() = %q, want %q", string(got), tt.want)
			}
		})
	}
}

func TestAutoSerializeNilReturnsNilSlice(t *testing.T) {
	got, err := Auto.Serialize(nil)
	if err != nil {
		t.Fatalf("Serialize(nil) error = %v", err)
	}
	if got != nil {
		t.Errorf("Serialize(nil) = %v, want nil", got)
	}
}

func TestAutoSerializeString(t *testing.T) {
	tests := []struct {
		name  string
		input any
		want  string
	}{
		{"nil", nil, ""},
		{"string", "hello", "hello"},
		{"bytes", []byte("raw"), "raw"},
		{"int", 42, "42"},
		{"map", map[string]any{"a": 1}, `{"a":1}`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Auto.SerializeString(tt.input)
			if err != nil {
				t.Fatalf("SerializeString() error = %v", err)
			}
			if got != tt.want {
				t.Errorf("SerializeString() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestAutoDeserialize(t *testing.T) {
	tests := []struct {
		name string
		input string
		want any
	}{
		{"json object", `{"a":1}`, map[string]any{"a": float64(1)}},
		{"json array", `[1,2]`, []any{float64(1), float64(2)}},
		{"json string", `"hello"`, "hello"},
		{"json number", `42`, float64(42)},
		{"json bool", `true`, true},
		{"json null", `null`, nil},
		{"plain text", `hello world`, "hello world"},
		{"path", `/foo/bar`, "/foo/bar"},
		{"date (false positive)", `2026-04-14`, "2026-04-14"},
		{"empty", ``, ""},
		{"whitespace json", `  {"a":1}`, map[string]any{"a": float64(1)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Auto.Deserialize([]byte(tt.input))
			if err != nil {
				t.Fatalf("Deserialize() error = %v", err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Deserialize() = %v (%T), want %v (%T)", got, got, tt.want, tt.want)
			}
		})
	}
}
