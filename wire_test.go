package wire

import "testing"

func TestByName(t *testing.T) {
	tests := []struct {
		name string
		want WireFormat
	}{
		{"auto", Auto},
		{"json", JSON},
		{"string", String},
		{"bytes", Bytes},
	}
	for _, tt := range tests {
		if got := ByName(tt.name); got != tt.want {
			t.Errorf("ByName(%q) = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestByNameUnknown(t *testing.T) {
	if got := ByName("unknown"); got != nil {
		t.Errorf("ByName(\"unknown\") = %v, want nil", got)
	}
}

func TestNames(t *testing.T) {
	tests := []struct {
		wf   WireFormat
		want string
	}{
		{Auto, "auto"},
		{JSON, "json"},
		{String, "string"},
		{Bytes, "bytes"},
	}
	for _, tt := range tests {
		if got := tt.wf.Name(); got != tt.want {
			t.Errorf("Name() = %q, want %q", got, tt.want)
		}
	}
}
