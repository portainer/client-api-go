# client-api
Swagger generated Client APIs in Golang

## Technical Details

### Pre-requisites

TBC

### Client Generation

Follow the steps below to generate Client APIs based on our `swagger.yaml` in Golang:

```sh
go mod init github.com/portainer/client-api-go
swagger generate client -f swagger.yaml -A portainer-client-api --principal portainer
go get -u ./...
```

The whole process is going to be automated by GitHub Actions.

### Release

Create a GitHub tag based on the Portainer repository's tag where the Swagger file is coming from.

## Utilising Client APIs

```
GOPRIVATE=github.com/portainer/client-api-go go get -u github.com/portainer/client-api-go/v2@v2.16.0
```


