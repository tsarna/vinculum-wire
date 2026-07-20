package wire

import "encoding/json"

// autoBytesFormat is auto's binary-preserving sibling: it decodes JSON the
// same way, but yields the undecoded []byte rather than a string when the
// payload isn't JSON.
//
// It exists for receivers carrying a mix of JSON and opaque binary on one
// stream. auto would stringify the binary case, which is lossless for
// storage but wrong in type — callers that want to hand the payload to
// something byte-oriented would have to convert it back.
//
// Serialization is identical to auto.
type autoBytesFormat struct{ autoFormat }

func (autoBytesFormat) Deserialize(b []byte) (any, error) {
	if !looksLikeJSON(string(b)) {
		return copyBytes(b), nil
	}

	var v any
	if err := json.Unmarshal(b, &v); err != nil {
		// JSON detection was a false positive (e.g. "2026-04-14" starts
		// with a digit but isn't valid JSON). Fall back to bytes.
		return copyBytes(b), nil
	}
	return v, nil
}

func (autoBytesFormat) Name() string { return "auto_bytes" }

// copyBytes returns a copy so the result doesn't alias a buffer the
// caller may reuse for the next message.
func copyBytes(b []byte) []byte {
	if b == nil {
		return nil
	}
	out := make([]byte, len(b))
	copy(out, b)
	return out
}
