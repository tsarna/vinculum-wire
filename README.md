# vinculum-wire

Pluggable wire-format system for converting between Go values and wire
representations (byte sequences). Used by
[vinculum](https://github.com/tsarna/vinculum) messaging clients and servers.

## Wire Formats

| Name | Serialize | Deserialize |
|------|-----------|-------------|
| `auto` | Strings and bytes verbatim; everything else JSON-encoded | Auto-detects JSON; falls back to string |
| `json` | All values JSON-encoded; bytes pass through | Strict JSON decode; errors on malformed input |
| `string` | Strings, bytes, numbers, bools to string form; rejects structured types | Returns string |
| `bytes` | Same as string | Returns `[]byte` |

All built-in formats pass `[]byte` input through unchanged on serialize.

## Usage

```go
import wire "github.com/tsarna/vinculum-wire"

// Use a built-in singleton
b, err := wire.JSON.Serialize(map[string]any{"status": "ok"})

// Look up by name
wf := wire.ByName("json")
val, err := wf.Deserialize(payload)

// Implement custom formats
type myFormat struct{}
func (myFormat) Serialize(v any) ([]byte, error)          { /* ... */ }
func (myFormat) SerializeString(v any) (string, error)    { /* ... */ }
func (myFormat) Deserialize(b []byte) (any, error)        { /* ... */ }
func (myFormat) Name() string                             { return "myformat" }
```

## Interface

```go
type WireFormat interface {
    Serialize(v any) ([]byte, error)
    SerializeString(v any) (string, error)
    Deserialize(b []byte) (any, error)
    Name() string
}
```

The interface works with `any`, not `cty.Value`, so this module has zero
external dependencies (stdlib only). The cty conversion layer lives in
vinculum itself.
