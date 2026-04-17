package wire

import "testing"

func TestLooksLikeJSON(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		// JSON values
		{`{"key":"val"}`, true},
		{`[1,2,3]`, true},
		{`"hello"`, true},
		{`42`, true},
		{`-1`, true},
		{`0.5`, true},
		{`true`, true},
		{`false`, true},
		{`null`, true},

		// Leading whitespace
		{`  {"key":"val"}`, true},
		{"\t\n42", true},

		// Not JSON
		{`hello world`, false},
		{`Hello`, false},
		{`/path/to/thing`, false},
		{``, false},
		{`   `, false},

		// False positives that looksLikeJSON accepts (by design):
		// "2026-04-14" starts with a digit. The auto format handles
		// these by catching json.Unmarshal errors and falling back.
		{`2026-04-14`, true},
	}
	for _, tt := range tests {
		if got := looksLikeJSON(tt.input); got != tt.want {
			t.Errorf("looksLikeJSON(%q) = %v, want %v", tt.input, got, tt.want)
		}
	}
}
