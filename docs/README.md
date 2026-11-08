# Public Usage

`GET /healthz` reports liveness, `/readyz` checks PostgreSQL/bootstrap schema,
and `/version` returns product/version/bootstrap stage. These unauthenticated
local-development endpoints do not certify training or data-processing readiness.
API default is loopback port 2056.

Go SDK only reads version. `VERSION` and OpenAPI info are `0.1.0`, not a release
claim. Bootstrap modules have no `/v3` suffix; future major paths follow actual
Go API versions. Nested SDK publication uses `sdk/go/vX.Y.Z`, separately from
the root product tag. No packages are published by local initialization.
