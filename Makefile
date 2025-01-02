BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
COMMIT := $(shell git log -1 --format='%H')
APPNAME := checkers

# don't override user values
ifeq (,$(VERSION))
  VERSION := $(shell git describe --exact-match 2>/dev/null)
  # if VERSION is empty, then populate it with branch's name and raw commit hash
  ifeq (,$(VERSION))
    VERSION := $(BRANCH)-$(COMMIT)
  endif
endif

# Update the ldflags with the app, client & server names
ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=$(APPNAME) \
	-X github.com/cosmos/cosmos-sdk/version.AppName=$(APPNAME)d \
	-X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
	-X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT)

BUILD_FLAGS := -ldflags '$(ldflags)'

##############
###  Test  ###
##############

test-unit:
	@echo Running unit tests...
	@go test -mod=readonly -v -timeout 30m ./...

test-race:
	@echo Running unit tests with race condition reporting...
	@go test -mod=readonly -v -race -timeout 30m ./...

test-cover:
	@echo Running unit tests and creating coverage report...
	@go test -mod=readonly -v -timeout 30m -coverprofile=$(COVER_FILE) -covermode=atomic ./...
	@go tool cover -html=$(COVER_FILE) -o $(COVER_HTML_FILE)
	@rm $(COVER_FILE)

bench:
	@echo Running unit tests with benchmarking...
	@go test -mod=readonly -v -timeout 30m -bench=. ./...

test: govet govulncheck test-unit

.PHONY: test test-unit test-race test-cover bench

#################
###  Install  ###
#################

all: install

install:
	@echo "--> ensure dependencies have not been modified"
	@go mod verify
	@echo "--> installing $(APPNAME)d"
	@go install $(BUILD_FLAGS) -mod=readonly ./cmd/$(APPNAME)d

.PHONY: all install

##################
###  Protobuf  ###
##################

# Use this proto-image if you do not want to use Ignite for generating proto files
protoVer=0.15.1
protoImageName=ghcr.io/cosmos/proto-builder:$(protoVer)
protoImage=$(DOCKER) run --rm -v $(CURDIR):/workspace --workdir /workspace $(protoImageName)

proto-gen:
	@echo "Generating protobuf files..."
	@ignite generate proto-go --yes

.PHONY: proto-gen

#################
###  Linting  ###
#################

golangci_lint_cmd=golangci-lint
golangci_version=v1.61.0

lint:
	@echo "--> Running linter"
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(golangci_version)
	@$(golangci_lint_cmd) run ./... --timeout 15m

lint-fix:
	@echo "--> Running linter and fixing issues"
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(golangci_version)
	@$(golangci_lint_cmd) run ./... --fix --timeout 15m

.PHONY: lint lint-fix

###################
### Development ###
###################

govet:
	@echo Running go vet...
	@go vet ./...

govulncheck:
	@echo Running govulncheck...
	@go install golang.org/x/vuln/cmd/govulncheck@latest
	@govulncheck ./...

mock-expected-keepers:
	@mockgen -source=x/checkers/types/expected_keepers.go \
		-package testutil \
		-destination=x/checkers/testutil/expected_keepers_mocks.go

gen-protoc-ts:
	@mkdir -p ./client/src/types/generated/
	@find protoc-gen/checkers/checkers -type f -printf '%p\n' | xargs -I {} protoc \
		--plugin="./scripts/node_modules/.bin/protoc-gen-ts_proto" \
		--ts_proto_out="./client/src/types/generated" \
		--proto_path="./protoc-gen" \
		--ts_proto_opt="esModuleInterop=true,forceLong=long,useOptionals=messages" \
		{}

.PHONY: govet govulncheck mock-expected-keepers gen-protoc-ts

###################
###    Build    ###
###################

build-linux:
	GOOS=linux GOARCH=amd64 go build -o ./build/checkersd-linux-amd64 ./cmd/checkersd/main.go
	GOOS=linux GOARCH=arm64 go build -o ./build/checkersd-linux-arm64 ./cmd/checkersd/main.go

do-checksum-linux:
	cd build && sha256sum \
		checkersd-linux-amd64 checkersd-linux-arm64 \
		> checkers-checksum-linux

build-linux-with-checksum: build-linux do-checksum-linux

build-darwin:
	GOOS=darwin GOARCH=amd64 go build -o ./build/checkersd-darwin-amd64 ./cmd/checkersd/main.go
	GOOS=darwin GOARCH=arm64 go build -o ./build/checkersd-darwin-arm64 ./cmd/checkersd/main.go

build-all: build-linux build-darwin

do-checksum-darwin:
	cd build && sha256sum \
		checkersd-darwin-amd64 checkersd-darwin-arm64 \
		> checkers-checksum-darwin

build-darwin-with-checksum: build-darwin do-checksum-darwin

build-with-checksum: build-linux-with-checksum build-darwin-with-checksum

.PHONY: build-linux do-checksum-linux build-linux-with-checksum build-darwin build-all do-checksum-darwin build-darwin-with-checksum build-with-checksum