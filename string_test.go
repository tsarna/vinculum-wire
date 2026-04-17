package wire

import "testing"

func TestStringSerialize(t *testing.T) {
	tests := []struct {
		name    string
		input   any
		want    string
		wantErr bool
	}{
		{"string", "hello", "hello", false},
		{"bytes", []byte("raw"), "raw", false},
		{"int", 42, "42", false},
		{"int64", int64(100), "100", false},
		{"float64", 3.14, "3.14", false},
		{"bool true", true, "true", false},
		{"bool false", false, "false", false},
		{"uint", uint(7), "7", false},
		{"nil", nil, "", true},
		{"map", map[string]any{}, "", true},
		{"slice", []any{1}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := String.Serialize(tt.input)
			if (err != nil) != tt.wantErr {
				t.Fatalf("Serialize() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && string(got) != tt.want {
				t.Errorf("Serialize() = %q, want %q", string(got), tt.want)
			}
		})
	}
}

func TestStringSerializeString(t *testing.T) {
	got, err := String.SerializeString("hello")
	if err != nil {
		t.Fatalf("SerializeString() error = %v", err)
	}
	if got != "hello" {
		t.Errorf("SerializeString() = %q, want %q", got, "hello")
	}
}

func TestStringDeserialize(t *testing.T) {
	input := []byte("hello world")
	got, err := String.Deserialize(input)
	if err != nil {
		t.Fatalf("Deserialize() error = %v", err)
	}
	s, ok := got.(string)
	if !ok {
		t.Fatalf("Deserialize() returned %T, want string", got)
	}
	if s != "hello world" {
		t.Errorf("Deserialize() = %q, want %q", s, "hello world")
	}
}

func TestStringDeserializeJSON(t *testing.T) {
	// String format always returns a string, even for JSON input.
	input := []byte(`{"a":1}`)
	got, err := String.Deserialize(input)
	if err != nil {
		t.Fatalf("Deserialize() error = %v", err)
	}
	s, ok := got.(string)
	if !ok {
		t.Fatalf("Deserialize() returned %T, want string", got)
	}
	if s != `{"a":1}` {
		t.Errorf("Deserialize() = %q, want %q", s, `{"a":1}`)
	}
}
