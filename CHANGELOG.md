# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.4.0] - 2026-07-20

### Added

- **`auto_bytes` wire format.** Decodes JSON exactly like `auto`, but returns the
  undecoded `[]byte` rather than a string when the payload isn't JSON — including when
  JSON detection turns out to be a false positive. Like `auto`, it never fails to decode.

  It exists for receivers carrying a mix of JSON and opaque binary on one stream, where
  `auto` would stringify the binary case. That is lossless for storage but wrong in type:
  a caller wanting to hand the payload to something byte-oriented would have to convert
  it back.

  As with `bytes`, the result is a copy, so it does not alias a read buffer the caller
  may reuse. Serialization is identical to `auto`.

## [0.3.0] - 2026-07-19

### Added

- `DecodeError` and `DecodeErrorHook` — a shared description of a failed deserialize
  (raw bytes, error, format name, best-effort topic, extracted fields, and a bag of
  per-client identity attributes) and the observer signature receivers invoke with it.
  Purely additive; nothing in this module produces or consumes them.

## [0.2.1] - 2026-06-26

### Fixed

- Re-release to repair the module checksum. The `v0.2.0` tag was inadvertently
  moved after publication, so its content no longer matched the hash recorded in
  the Go checksum database, breaking `go build`/`go mod` for downstream consumers.
  `v0.2.0` should be considered poisoned; use `v0.2.1` or later instead. No API
  changes from `v0.2.0`.

## [0.2.0] - 2026-05-25

### Changed

- **License changed to Apache 2.0** — the project is now licensed under the Apache License, Version 2.0, replacing the previous BSD 2-Clause license. A `NOTICE` file has been added per the Apache license requirements.

## [0.1.0] - 2026-04-17

### Added

- **`WireFormat` interface** — pluggable wire-format system for converting between Go values and wire representations (byte sequences).
- **Built-in formats** — `auto`, `json`, `string`, and `bytes`, available as singletons and via `ByName` lookup.
