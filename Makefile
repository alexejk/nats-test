
APP_BUILD_DATE=`date +%Y-%m-%d`
APP_BUILD_TIME=`date +%H:%M`
GO_FLAGS= CGO_ENABLED=0
GO_LDFLAGS= -ldflags=""
GO_BUILD_CMD=$(GO_FLAGS) go build $(GO_LDFLAGS)
BUILD_DIR=build
BINARY_NAME=nats-test

all: clean test lint build

lint:
	@echo "Linting code..."
	@go vet `go list ./... | grep -v $(MOCK_DIR)`

test:
	@echo "Running tests..."
	@go test `go list ./... | grep -v $(MOCK_DIR)`

pre-build:
	@mkdir -p $(BUILD_DIR)

build-linux: pre-build
	@echo "Building Linux binary..."
	GOOS=linux GOARCH=amd64 $(GO_BUILD_CMD) -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64

build-osx: pre-build
	@echo "Building OSX binary..."
	GOOS=darwin GOARCH=amd64 $(GO_BUILD_CMD) -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64

clean:
	@echo "Cleanup.."
	@rm -Rf $(BUILD_DIR)