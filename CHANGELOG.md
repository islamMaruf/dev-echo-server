# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.0] - 2025-11-15

### Added

- Initial release of Dev Echo Server
- HTTP echo server for development and testing
- Security headers middleware (Helmet.js-like)
- JSON request/response logging with file rotation
- Error handling middleware with panic recovery
- Environment-based configuration (.env support)
- Docker support with multi-stage builds
- Docker Compose configuration
- Interactive run script (run.sh) with automatic fallback
- Makefile with common development tasks
- Comprehensive documentation
- Support for both `docker compose` and `docker-compose`
- Automatic .env file creation from .env.example
- Health check endpoint
- Request ID generation for tracking
- Catch-all route that echoes request body
- Clean project structure with separated concerns

### Features

- Port configuration via PORT environment variable
- Development/Production mode via NODE_ENV
- Background process support
- Request body logging and mirroring
- Graceful server startup with timeouts
- Log directory auto-creation

[1.0.0]: https://github.com/islamMaruf/dev-echo-server/releases/tag/v1.0.0
