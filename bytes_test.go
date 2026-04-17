package wire

import "testing"

func TestBytesSerialize(t *testing.T) {
	tests := []struct {
		name    string
		input   any
		want    string
		wantErr bool
	}{
		{"string", "hello", "hello", false},
		{"bytes", []byte("raw"), "raw", false},
		{"int", 42, "42", false},
		{"bool", true, "true", false},
		{"nil", nil, "", true},
		{"map", map[string]any{}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Bytes.Serialize(tt.input)
			if (err != nil) != tt.wantErr {
				t.Fatalf("Serialize() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && string(got) != tt.want {
				t.Errorf("Serialize() = %q, want %q", string(got), tt.want)
			}
		})
	}
}

func TestBytesSerializeString(t *testing.T) {
	got, err := Bytes.SerializeString("hello")
	if err != nil {
		t.Fatalf("SerializeString() error = %v", err)
	}
	if got != "hello" {
		t.Errorf("SerializeString() = %q, want %q", got, "hello")
	}
}

func TestBytesDeserialize(t *testing.T) {
	input := []byte("hello world")
	got, err := Bytes.Deserialize(input)
	if err != nil {
		t.Fatalf("Deserialize() error = %v", err)
	}
	b, ok := got.([]byte)
	if !ok {
		t.Fatalf("Deserialize() returned %T, want []byte", got)
	}
	if string(b) != "hello world" {
		t.Errorf("Deserialize() = %q, want %q", string(b), "hello world")
	}
}

func TestBytesDeserializeReturnsCopy(t *testing.T) {
	input := []byte("original")
	got, _ := Bytes.Deserialize(input)
	b := got.([]byte)

	// Mutate the output — should not affect input.
	b[0] = 'X'
	if input[0] == 'X' {
		t.Error("Deserialize() returned a slice sharing the input backing array")
	}
}

func TestBytesDeserializeJSON(t *testing.T) {
	// Bytes format always returns []byte, even for JSON input.
	input := []byte(`{"a":1}`)
	got, err := Bytes.Deserialize(input)
	if err != nil {
		t.Fatalf("Deserialize() error = %v", err)
	}
	b, ok := got.([]byte)
	if !ok {
		t.Fatalf("Deserialize() returned %T, want []byte", got)
	}
	if string(b) != `{"a":1}` {
		t.Errorf("Deserialize() = %q, want %q", string(b), `{"a":1}`)
	}
}
