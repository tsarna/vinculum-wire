// Package wire defines a pluggable wire-format system for converting
// between Go values and wire representations (byte sequences).
package wire

// WireFormat converts between Go values and wire representations.
//
// Implementations must be safe for concurrent use.
type WireFormat interface {
	// Serialize converts a Go value to its wire representation.
	//
	// Accepted input types depend on the implementation:
	//   - json: any JSON-marshalable value (maps, slices, scalars, etc.)
	//   - string/bytes: string, []byte, numbers, bools
	//   - auto: string and []byte verbatim; everything else JSON-encoded
	//
	// []byte input passes through unchanged in all built-in wire formats.
	Serialize(v any) ([]byte, error)

	// SerializeString is like Serialize but returns a string.
	//
	// Implementations may optimize this (e.g. JSON avoids the []byte →
	// string copy). The default behavior is equivalent to
	// string(Serialize(v)).
	SerializeString(v any) (string, error)

	// Deserialize converts a wire payload to a Go value.
	//
	// Return types depend on the implementation:
	//   - json: map[string]any, []any, string, float64, bool, or nil
	//   - string: string
	//   - bytes: []byte
	//   - auto: JSON-detected → natural Go types; otherwise string
	Deserialize(b []byte) (any, error)

	// Name returns the wire format's identifier (e.g. "json", "auto").
	Name() string
}

// Built-in singletons.
var (
	Auto   WireFormat = autoFormat{}
	JSON   WireFormat = jsonFormat{}
	String WireFormat = stringFormat{}
	Bytes  WireFormat = bytesFormat{}
)

var builtins = map[string]WireFormat{
	"auto":   Auto,
	"json":   JSON,
	"string": String,
	"bytes":  Bytes,
}

// ByName returns a built-in WireFormat by name, or nil if unknown.
func ByName(name string) WireFormat {
	return builtins[name]
}
