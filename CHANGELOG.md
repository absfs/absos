# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Comprehensive GoDoc comments for all exported types and interfaces
- Context support (`context.Context`) for all I/O operations
- Custom error types (`BucketError`, `ObjectError`) with error wrapping
- Named parameters for better API clarity (e.g., `key string` instead of `string`)
- In-memory reference implementation in `examples/memory`
- GitHub Actions CI/CD workflow with:
  - Multi-platform testing (Ubuntu, macOS, Windows)
  - Go version matrix testing (1.21, 1.22, 1.23)
  - Linting with golangci-lint
  - Code formatting checks
  - Go vet analysis
- Comprehensive test suite for error types and memory implementation
- CONTRIBUTING.md with contribution guidelines
- .gitignore for Go projects
- .golangci.yml for consistent linting
- Enhanced README with:
  - Usage examples with context
  - Architecture documentation
  - Installation instructions
  - Feature list
  - Badges (Go Reference, Go Report Card, License)
- CHANGELOG.md for tracking version history

### Changed
- Fixed Go version in go.mod from invalid `1.25.4` to `1.21`
- Updated all interface methods to accept `context.Context` as first parameter
- Enhanced documentation throughout the codebase

### Fixed
- Corrected invalid Go version specification

## [0.1.0] - Initial Release

### Added
- Core interfaces: `ObjectStore`, `Bucket`, `Object`, `ObjectHeader`
- Support interfaces: `Owner`, `Page`, `SSE`
- Basic object storage operations
- Pagination support for listing objects
- Server-side encryption configuration
- MIT License

[Unreleased]: https://github.com/absfs/absos/compare/v0.1.0...HEAD
[0.1.0]: https://github.com/absfs/absos/releases/tag/v0.1.0
