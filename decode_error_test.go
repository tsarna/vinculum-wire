package wire

import "testing"

// TestIsReservedAttr pins the set of names a receiver must not use as an Attrs
// key. Each one is the eval-context name of a fixed DecodeError field, so a
// receiver that reuses it loses its value to the consumer's collision check —
// which is how vinculum-mqtt's Attrs["topic"] went unnoticed.
func TestIsReservedAttr(t *testing.T) {
	for _, key := range []string{"raw", "error", "wire_format", "topic", "fields"} {
		if !IsReservedAttr(key) {
			t.Errorf("IsReservedAttr(%q) = false, want true", key)
		}
	}

	// Transport-named identifiers are exactly what Attrs is for, and none of
	// them collide.
	for _, key := range []string{
		"mqtt_topic", "routing_key", "exchange", "queue",
		"stream", "entry_id", "group", "consumer",
		"partition", "offset",
	} {
		if IsReservedAttr(key) {
			t.Errorf("IsReservedAttr(%q) = true, want false", key)
		}
	}
}
