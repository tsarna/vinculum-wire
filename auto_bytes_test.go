package wire

import (
	"reflect"
	"testing"
)

func TestAutoBytesDeserializeJSONMatchesAuto(t *testing.T) {
	// The JSON paths must behave identically to auto.
	for _, in := range []string{`{"a":1}`, `[1,2,3]`, `"hello"`, `42`, `true`, `null`} {
		t.Run(in, func(t *testing.T) {
			got, err := AutoBytes.Deserialize([]byte(in))
			if err != nil {
				t.Fatalf("AutoBytes.Deserialize(%s) error = %v", in, err)
			}
			want, err := Auto.Deserialize([]byte(in))
			if err != nil {
				t.Fatalf("Auto.Deserialize(%s) error = %v", in, err)
			}
			if !reflect.DeepEqual(got, want) {
				t.Errorf("AutoBytes.Deserialize(%s) = %#v, want %#v (same as auto)", in, got, want)
			}
		})
	}
}

func TestAutoBytesDeserializeNonJSON(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{"plain text", "not json {{"},
		// Starts with a digit so looksLikeJSON accepts it, but it does not
		// parse — the same false positive auto falls back on.
		{"json false positive", "2026-04-14"},
		{"empty", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AutoBytes.Deserialize([]byte(tt.input))
			if err != nil {
				t.Fatalf("Deserialize() error = %v, want nil (auto_bytes never fails)", err)
			}
			b, ok := got.([]byte)
			if !ok {
				t.Fatalf("Deserialize() = %T, want []byte", got)
			}
			if string(b) != tt.input {
				t.Errorf("Deserialize() = %q, want %q", b, tt.input)
			}
		})
	}
}

func TestAutoBytesDiffersFromAutoOnNonJSON(t *testing.T) {
	// This difference is the entire reason the format exists.
	raw := []byte("not json {{")

	ab, err := AutoBytes.Deserialize(raw)
	if err != nil {
		t.Fatalf("AutoBytes.Deserialize() error = %v", err)
	}
	if _, ok := ab.([]byte); !ok {
		t.Errorf("AutoBytes.Deserialize() = %T, want []byte", ab)
	}

	a, err := Auto.Deserialize(raw)
	if err != nil {
		t.Fatalf("Auto.Deserialize() error = %v", err)
	}
	if _, ok := a.(string); !ok {
		t.Errorf("Auto.Deserialize() = %T, want string", a)
	}
}

func TestAutoBytesPreservesBinary(t *testing.T) {
	raw := []byte{0xff, 0xfe, 0x00, 0x01, 0x80}
	got, err := AutoBytes.Deserialize(raw)
	if err != nil {
		t.Fatalf("Deserialize() error = %v", err)
	}
	if !reflect.DeepEqual(got, raw) {
		t.Errorf("Deserialize() = %#v, want %#v (arbitrary binary must survive)", got, raw)
	}
}

func TestAutoBytesDoesNotAliasInput(t *testing.T) {
	// Receivers commonly reuse one read buffer across messages.
	buf := []byte("not json")
	got, err := AutoBytes.Deserialize(buf)
	if err != nil {
		t.Fatalf("Deserialize() error = %v", err)
	}
	buf[0] = 'X'
	if string(got.([]byte)) != "not json" {
		t.Errorf("Deserialize() result aliases the caller's buffer: got %q", got)
	}
}

func TestAutoBytesSerializeMatchesAuto(t *testing.T) {
	inputs := []any{nil, "str", []byte("raw"), map[string]any{"a": 1}, 42, true}
	for _, in := range inputs {
		got, err := AutoBytes.Serialize(in)
		if err != nil {
			t.Fatalf("AutoBytes.Serialize(%#v) error = %v", in, err)
		}
		want, err := Auto.Serialize(in)
		if err != nil {
			t.Fatalf("Auto.Serialize(%#v) error = %v", in, err)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("AutoBytes.Serialize(%#v) = %q, want %q (same as auto)", in, got, want)
		}

		gotS, err := AutoBytes.SerializeString(in)
		if err != nil {
			t.Fatalf("AutoBytes.SerializeString(%#v) error = %v", in, err)
		}
		wantS, err := Auto.SerializeString(in)
		if err != nil {
			t.Fatalf("Auto.SerializeString(%#v) error = %v", in, err)
		}
		if gotS != wantS {
			t.Errorf("AutoBytes.SerializeString(%#v) = %q, want %q", in, gotS, wantS)
		}
	}
}

func TestAutoBytesByName(t *testing.T) {
	wf := ByName("auto_bytes")
	if wf == nil {
		t.Fatal(`ByName("auto_bytes") = nil, want the AutoBytes singleton`)
	}
	if wf.Name() != "auto_bytes" {
		t.Errorf("Name() = %q, want %q", wf.Name(), "auto_bytes")
	}
}
