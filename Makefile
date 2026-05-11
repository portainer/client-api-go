VERSION := 2.31.2

.PHONY: help generate-client

default: help

# Inspired from https://dwmkerr.com/makefile-help-command/
help:
	@grep -E '^[a-zA-Z0-9 -]+:.*#'  Makefile | sort | while read -r l; do printf "\033[1;32m$$(echo $$l | cut -f 1 -d':')\033[00m:$$(echo $$l | cut -f 2- -d'#')\n"; done

generate-client:
	@echo "Generating client for Portainer API version $(VERSION)"
	curl -o swagger.yaml https://api.swaggerhub.com/apis/portainer/portainer-ee/$(VERSION)/swagger.yaml
	# --skip-validation: the upstream Portainer EE swagger spec contains OpenAPI 2.0
	# violations that abort `swagger generate` if validation runs. Remove this flag
	# once the upstream spec has been cleaned up.
	go run -modfile=tools/go.mod github.com/go-swagger/go-swagger/cmd/swagger generate client -f swagger.yaml -A portainer-client-api --principal portainer --skip-validation --target=pkg --client-package=client --model-package=models
	@echo "Client generation complete"

test:
	go test ./...