# set makefile echo back
ifdef VERBOSE
	V :=
else
	V := @
endif

.PHONY: build
## build : Build binary(arrow-study)
build:
	@echo "Building..."
	$(V)go build

.PHONY: lint
## lint : Lint codespace
lint:
	$(V)golangci-lint help linters

.PHONY: fmt
## fmt : Format all code
fmt:
	$(V)gofmt -w .

.PHONY: test
## test : Run test
test:
	$(V)go test -v .

.PHONY: help
## help : Print help message
help: Makefile
	@sed -n 's/^##//p' $< | awk 'BEGIN {FS = ":"} {printf "\033[36m%-13s\033[0m %s\n", $$1, $$2}'

.PHONY: bitmap_of_array
## bitmap_of_array : Build bitmap_of_array
bitmap_of_array:
	$(V)go build bin/bitmap_of_array/bitmap_of_array.go