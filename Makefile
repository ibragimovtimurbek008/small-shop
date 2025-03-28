.PHONY: build
build:
	go build -o small_shop ./cmd/

.PHONY: local_auth
local_auth:
	APPLICATION_ADDR=:8080 \
	./small_shop auth

.PHONY: vendor
vendor:
	go mod tidy
	go mod vendor

%:
	@:

