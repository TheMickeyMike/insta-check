GIT_VERSION := $(shell git describe --abbrev=4 --dirty --always --tags)
LDFLAGS :=-ldflags "-X main.version=$(GIT_VERSION)"
GO_TEST :=env GOTRACEBACK=all GO111MODULE=on go test -v $(LDFLAGS)
GO_BUILD :=env GO111MODULE=on go build $(LDFLAGS)
GOINSTALL :=go install

.PHONY: all

all: clean integration-test build

test:
	$(GO_TEST) ./...

integration-test:
	$(GO_TEST) -tags integration ./...

build:
	$(GO_BUILD) -o ./build/app cmd/instachecker/*

mocks:
	$(GOINSTALL) github.com/golang/mock/mockgen
	mockgen -source=pkg/client/http.go -destination=pkg/client/mocks/http_mock.go -package=mocks

clean:
	@rm -rf build
