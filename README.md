# DIDACTOVELA Open

Public integration surface for data and learned-policy lifecycle management.

| Repository | Scope | License |
| --- | --- | --- |
| [didactovela-open](https://github.com/axisvela/didactovela-open) | Public contracts, SDK, examples and approved core binaries | Apache-2.0 |
| [didactovela](https://github.com/axisvela/didactovela) | Go/PostgreSQL core | AGPL-3.0-only |
| [didactovela-ee](https://github.com/axisvela/didactovela-ee) | Enterprise extensions and internal product documentation | Enterprise |

Unreleased `0.1.0`, bootstrap maturity C. Only operational endpoints and a Go
version client exist; capture, dataset, annotation and training APIs do not.
No datasets, weights, trainers or binaries are included. Target GitHub addresses
are not package-registry or release availability claims.

## Quickstart

Start the core using its README (API `127.0.0.1:2056`, PostgreSQL loopback `2057`).
From `sdk/go/`, with Go 1.25+:

```powershell
go test ./...
go vet ./...
go run ./examples/version http://127.0.0.1:2056
```

The source-tree example only reads version from a running core. It is not a
published-package check. Operational contract:
[contracts/bootstrap.openapi.json](contracts/bootstrap.openapi.json).
Public usage: [docs/README.md](docs/README.md).

Training outputs remain candidates, not authorization policy or automatic
production deployment. DIDACTOVELA does not replace inference, simulation,
world-state authority or independent assurance.

Apache-2.0 covers original public source, not redistributed AGPL core binaries.
Future binary bundles must include matching source, license notices, SBOM,
provenance and checksums. No binaries are currently published.
