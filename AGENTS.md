# Agent Guidelines for InfooScreen

Multi-service screen sharing application with TypeScript/WebSocket rendezvous server and Go web server deployed on Vercel.

## Build, Lint, and Test Commands

**TypeScript (screensy-rendezvous):**
- Build: `cd screensy-rendezvous && npx tsc -p tsconfig.json`
- Install deps: `cd screensy-rendezvous && npm install`
- No tests configured

**Go (screensy-website):**
- Build: `cd screensy-website && go build .`
- Run: `cd screensy-website && go run main.go`
- Test: `cd screensy-website && go test ./...`
- Single test: `cd screensy-website && go test -run TestName`
- Lint/Format: `cd screensy-website && go vet ./... && go fmt ./...`

**Docker (Full stack):**
- Build all: `docker-compose build`
- Run all: `docker-compose up -d`
- Run single service: `docker-compose up <service-name>`

## Code Style Guidelines

**TypeScript:** Strict mode enabled. 4-space indent. `camelCase` variables/functions, `PascalCase` interfaces. Explicit imports. JSDoc comments explaining *why*.

**Go:** `go fmt` standard formatting. `camelCase` unexported, `PascalCase` exported. Explicit error handling with `if err != nil`. Standard comments.

**General:** No tests exist yet. Focus on WebRTC signaling and i18n. Docker-first deployment with Caddy reverse proxy.

## Cursor/Copilot Rules

No specific Cursor or Copilot rules files found.
