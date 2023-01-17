CGO_ENABLED:=CGO_ENABLED=0

.PHONY: test
test:
	@$(CGO_ENABLED) go test ./... -v -cover -coverprofile=cov.out

.PHONY: vendor
vendor:
	@$(CGO_ENABLED) go mod vendor

.PHONY: clean
clean:
	@go clean
	@rm -rf ./bin/*
