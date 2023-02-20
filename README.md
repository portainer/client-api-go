# client-api
Swagger generated Client APIs in Golang

## Technical Details

### Pre-requisites

TBC

### Client Generation

Follow the steps below to generate Client APIs based on our `swagger.yaml` in Golang:

```sh
go mod init github.com/portainer/client-api-go
swagger generate client -f swagger.yaml -A portainer-client-api --principal portainer --skip-validation
go get -u ./...
```

Then, create a GitHub tag based on the Portainer repository's tag where the Swagger file is coming from.

```sh
git tag -a v1.24.0 -m "Portainer v1.24.0"
git push origin v1.24.0
```

The whole process is going to be automated by GitHub Actions.

### Release

Create a GitHub tag based on the Portainer repository's tag where the Swagger file is coming from.

## Utilising Client APIs

TBC
