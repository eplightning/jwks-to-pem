BINARY := jwks-to-pem
PLATFORMS := linux-amd64 linux-arm64 windows-amd64 darwin-amd64
VERSION := 0.1
GO_BUILD_FLAGS =

os = $(word 1,$(subst -, ,$(firstword $@)))
arch = $(word 2,$(subst -, ,$(firstword $@)))

all: build

.PHONY: build
build:
	CGO_ENABLED=0 go build $(GO_BUILD_FLAGS) -o $(BINARY)

.PHONY: $(PLATFORMS)
$(PLATFORMS):
	mkdir -p build
	CGO_ENABLED=0 GOOS=$(os) GOARCH=$(arch) go build $(GO_BUILD_FLAGS) -o build/$(BINARY)-$(VERSION)-$(os)-$(arch)

.PHONY: release
release: $(PLATFORMS)

.PHONY: clean
clean:
	rm -rf build
	rm -f $(BINARY)
