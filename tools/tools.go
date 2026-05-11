//go:build tools

// Package tools pins build-time tools used to regenerate the swagger
// client. It lives in its own module (tools/go.mod) so the tool's
// transitive dependencies do not bleed into the SDK's runtime
// dependency graph or force a Go-version bump on consumers.
//
// The blank imports below exist solely to keep `go mod tidy` (run
// inside tools/) from pruning the tool modules from go.mod. The
// `tools` build tag ensures this file is never compiled into anything
// useful — it is just a manifest.
//
// Invoke a tool via the Makefile, which does the equivalent of:
//
//	cd tools && go run github.com/go-swagger/go-swagger/cmd/swagger ...
package tools

import (
	_ "github.com/go-swagger/go-swagger/cmd/swagger"
)
