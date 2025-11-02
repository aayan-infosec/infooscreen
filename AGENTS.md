# Agent Guidelines for InfooScreen.github.io

This document outlines the conventions and commands for agentic coding in this repository.

## 1. Build, Lint, and Test Commands

*   **TypeScript (screensy-rendezvous, screensy-website):**
    *   Build: `npx tsc -p <project-directory>/tsconfig.json` (e.g., `npx tsc -p screensy-rendezvous/tsconfig.json`)
    *   Run all tests: No explicit test runner configured.
*   **Go (screensy-website):**
    *   Build: `go build ./...` (run from `screensy-website/`)
    *   Run all tests: `go test ./...` (run from `screensy-website/`)
    *   Run a single test: `go test -run <TestName> ./...` (run from `screensy-website/`)
    *   Lint/Format: `go vet ./...` and `go fmt ./...` (run from `screensy-website/`)

## 2. Code Style Guidelines

*   **Imports:** Use explicit imports. For TypeScript, `import * as ...` is common.
*   **Formatting:**
    *   **TypeScript:** 4-space indentation. Use `camelCase` for variables and functions. Interfaces are preferred for type definitions.
    *   **Go:** Use `go fmt` for standard formatting (tabs for indentation). `camelCase` for unexported identifiers, `PascalCase` for exported.
*   **Types:** Leverage TypeScript's strict type checking (`strict: true` in `tsconfig.json`). Go is statically typed.
*   **Naming Conventions:**
    *   **TypeScript:** `camelCase` for variables, functions, and methods. `PascalCase` for interfaces and classes.
    *   **Go:** `camelCase` for unexported variables/functions, `PascalCase` for exported.
*   **Error Handling:**
    *   **TypeScript:** Use `try...catch` for runtime errors (e.g., JSON parsing). Validate inputs explicitly.
    *   **Go:** Handle errors explicitly using `if err != nil { ... }`. Use `log.Fatal` for critical errors and `panic` for unrecoverable initialization issues.
*   **Comments:** Use JSDoc for TypeScript and standard Go comments. Explain *why* rather than *what*.

## 3. Cursor/Copilot Rules

No specific Cursor or Copilot rules files were found in this repository.
