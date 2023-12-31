dependencies:
	$(GOGET) -u google.golang.org/grpc
	$(GOGET) -u github.com/grpc-ecosystem/grpc-gateway
	$(GOGET) -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	$(GOGET) -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	$(GOGET) -u github.com/golang/protobuf/protoc-gen-go

build-proto-go:
	find ./proto/messages -name "*.proto" 2>/dev/null | xargs realpath | xargs -I {} protoc $(PROTOC_IMPORT_PATH) --go_out=plugins=grpc:$(BUILD_PROTO_DIRECTORY) {}
	find ./proto/services -name "*.proto" 2>/dev/null | xargs realpath | xargs -I {} protoc $(PROTOC_IMPORT_PATH) --grpc-gateway_out=logtostderr=true:$(BUILD_PROTO_DIRECTORY) --go_out=plugins=grpc:$(BUILD_PROTO_DIRECTORY) {}

rate-limiter:
	@echo "Building Rate limiter"
	$(GOBUILD) ./...
	$(GOBUILD) -o $(BINARY_NAME) -v

docker-build:
	docker build -t $(DOCKER_IMAGE_NAME):latest .

all: build-proto-go

all-by-docker:
	docker run --rm --name proto-compiler -it --user $$(id -u):$$(id -g) $(DOCKER_VOLUMES) $(DOCKER_IMAGE_NAME) all

clean:
	@echo "Cleanning..."
	-rm -f $(BINARY_NAME)
	-rm -f $(BINARY_UNIX)
	-rm -f $(BINARY_SPEC)
	-rm -rf build
	-rm -rf dist
	@echo "Done cleanning."

deployment:
	docker-compose -f ./deployment/compose/docker-compose.yml up -d


GOCMD=go
GOMOD=GO111MODULE=on $(GOCMD) mod
GOGET=GO111MODULE=on $(GOCMD) "get"

BUILD_DIRECTORY=build
BUILD_PATH=git.bale.ir/
PROJECT_PATH=git.bale.ir/rate_limiter/

BUILD_PROTO_DIRECTORY=../
GOOGLE_APIS_DIR="$$(find $(GOPATH) -wholename "*github.com/grpc-ecosystem/grpc-gateway*/third_party/googleapis" 2>/dev/null | head -n 1)"
PROTOC_IMPORT_PATH=-I${GOOGLE_APIS_DIR} -I $$PWD/proto -I/usr/local/include
GOBUILD=$(GOCMD) build
BINARY_NAME=rate_limiter
BINARY_UNIX=$(BINARY_NAME)_unix
BINARY_SPEC=$(BINARY_NAME).specs
DOCKER_VOLUMES=-v $$PWD:/go/src/$(PROJECT_PATH)
DOCKER_IMAGE_NAME=git.bale.ir:8000/rate_limiter
