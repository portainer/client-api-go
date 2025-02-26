# client-api

Swagger generated Client APIs in Golang

## Technical Details

### Pre-requisites

### Install Swagger CLI

```sh
go install github.com/go-swagger/go-swagger/cmd/swagger@latest
```

### Download Swagger file from SwaggerHub

```sh
curl -o swagger.yaml https://raw.githubusercontent.com/portainer/portainer/master/api/http/swagger.yaml
```

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

```sh
GOPRIVATE=github.com/portainer/client-api-go go get -u github.com/portainer/client-api-go/v2@v2.16.0
```
