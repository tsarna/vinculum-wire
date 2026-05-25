# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.2.0] - 2026-05-25

### Changed

- **License changed to Apache 2.0** — the project is now licensed under the Apache License, Version 2.0, replacing the previous BSD 2-Clause license. A `NOTICE` file has been added per the Apache license requirements.

## [0.1.0] - 2026-04-17

### Added

- **`WireFormat` interface** — pluggable wire-format system for converting between Go values and wire representations (byte sequences).
- **Built-in formats** — `auto`, `json`, `string`, and `bytes`, available as singletons and via `ByName` lookup.
