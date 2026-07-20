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
	Attrs map[string]string
}

// DecodeErrorHook observes a deserialize failure.
//
// It is purely an observer: the receiver treats the message as failed
// regardless of what the hook does, and a hook that panics or errors must
// not change that outcome. A nil hook means "no observer configured".
type DecodeErrorHook func(ctx context.Context, e DecodeError)
