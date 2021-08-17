install:
	which swagger || GO11MODULE=off go get -u github.com/go-swagger/go-swagger/cmd/swagger

swagger: install
	swagger generate spec -o ./swagger.yaml --scan-models
