# Repository Guidelines

## Project Structure & Module Organization
- `cmd/server`: entrypoint that builds the `livekit-server` binary.
- `pkg`: core server code (RTC pipeline, signaling, auth, services) plus generated wiring; treat it as the primary place for new modules.
- `test`: integration and load scenarios; mirrors package layouts for black-box coverage.
- `tools` and `magefile*.go`: developer utilities and Mage tasks; avoid editing generated outputs under `pkg/service/wire_gen.go` unless regenerating.
- `config-sample.yaml`: starting point for local configuration; copy and adjust rather than editing in place.

## Build, Test, and Development Commands
- `./bootstrap.sh`: install Go toolchain helpers (Mage, codegen deps); run once per environment or after tool upgrades.
- `mage` or `mage Build`: compile the server to `bin/livekit-server` after regenerating wiring if needed.
- `mage BuildLinux`: cross-compile for Linux; set `GOARCH` when targeting a specific architecture.
- `mage Generate`: run `go generate` and refresh DI wiring; commit the resulting files with your change.
- `mage Clean`: remove build artifacts and checksum files.
- Run locally after build: `bin/livekit-server --config config-sample.yaml`.

## Coding Style & Naming Conventions
- Go 1.23+; keep GOPATH/bin on PATH so mage-installed tools are found.
- Format with `gofmt`/`goimports` and keep imports grouped; avoid hand-editing generated files.
- Package names stay short and lowercase; exported symbols need GoDoc-style comments.
- Prefer small interfaces and explicit context propagation; guard shared state with sync primitives already used in neighboring code.
- Avoid introducing new dependencies unless necessary and aligned with existing patterns.

## Testing Guidelines
- `mage Test`: fast suite (`go test -short ./... -count=1`) for unit coverage.
- `mage TestAll`: full suite with integration paths (`go test ./... -count=1 -timeout=4m -v`).
- Place tests in `*_test.go` alongside the code; name cases `TestXxx` and table-driven where possible.
- Keep tests deterministic and avoid network-bound external calls; use fakes/mocks already present in the package.

## Commit & Pull Request Guidelines
- Use concise, imperative commit subjects (e.g., “Add SFU track stats guard”); keep scope focused.
- Describe changes, risks, and validation steps in PRs; link issues and note config/schema impacts.
- Include logs or sample commands when behavior changes; add screenshots only if a user-facing change exists.
- Avoid force-pushes on shared branches and do not rewrite history already reviewed.
