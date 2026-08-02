package wire

import "context"

// DecodeError describes a failure to deserialize an inbound message body.
//
// Receivers construct one when WireFormat.Deserialize returns an error and
// pass it to their configured DecodeErrorHook. It carries everything needed
// to diagnose or reroute the message: the undecoded bytes, the error, and
// whatever identifying information the receiver had in hand at the time.
type DecodeError struct {
	// Raw is the undecoded message body.
	Raw []byte

	// Err is the error returned by Deserialize.
	Err error

	// Format is the configured wire format's Name (e.g. "json").
	Format string

	// Topic is the vinculum topic the message would have been delivered
	// on, best-effort. Receivers whose topic derives from the decoded
	// message fall back to the transport-level source name (routing key,
	// channel, stream, ...).
	Topic string

	// Fields holds any fields already extracted from the routing key,
	// headers, or equivalent. May be nil.
	Fields map[string]string

	// Attrs holds per-client identifying information — for example
	// routing_key and exchange for rabbitmq, partition and offset for
	// kafka, stream and entry_id for redis streams. Keys become
	// attributes on the hook's eval context. May be nil.
	//
	// Name a key after the transport, not after the concept it stands in
	// for: mqtt_topic rather than topic, so the key says which layer it
	// describes and cannot collide with a field above.
	//
	// The names IsReservedAttr reports are taken by the fixed fields above.
	// A consumer is expected to drop a colliding key rather than let a client
	// shadow Topic or Raw, so a collision silently loses the value.
	Attrs map[string]string
}

// reservedAttrs are the names the fixed DecodeError fields occupy once a
// consumer projects them onto an eval context.
var reservedAttrs = map[string]bool{
	"raw":         true, // Raw
	"error":       true, // Err
	"wire_format": true, // Format
	"topic":       true, // Topic
	"fields":      true, // Fields
}

// IsReservedAttr reports whether an Attrs key collides with one of the fixed
// DecodeError fields. Receivers must not use a reserved key: a consumer drops
// it rather than let a client shadow Topic or Raw, so the value is lost.
//
// Name a transport identifier after its transport — mqtt_topic, routing_key,
// stream — and no collision arises in the first place.
func IsReservedAttr(key string) bool {
	return reservedAttrs[key]
}

// DecodeErrorHook observes a deserialize failure.
//
// It is purely an observer: the receiver treats the message as failed
// regardless of what the hook does, and a hook that panics or errors must
// not change that outcome. A nil hook means "no observer configured".
type DecodeErrorHook func(ctx context.Context, e DecodeError)
