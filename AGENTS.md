# DIDACTOVELA Open Rules

- Use installed Superpowers, following `superpowers/README.md`; report if missing.
- Keep internal architecture/design/plans in EE. This tree owns public API, SDK,
  examples, usage and approved distributions only.
- Apache SDKs must not import AGPL core or cross-product internals.
- From `sdk/go/`: `go test ./...`, `go vet ./...`, `go build ./...`.
  From root: `python -m unittest discover -s tests -v`.
- No secrets, customer data, weights or arbitrary checkpoint executables in samples.
- Version contracts before consumers; no claims of nonexistent learning APIs.
- Versions X.Y.Z; core/open align, EE independently declares core compatibility.
- No implicit commits, tags, publication or push. Service tests run in background.
