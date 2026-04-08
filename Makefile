GO = $(shell which go 2>/dev/null)

APP             := linear
SRC_DIR         := cmd/linear-cli
VERSION         ?= v0.1.0
LDFLAGS         := -ldflags "-X main.AppVersion=$(VERSION)"
PREFIX          ?= /usr/local

.PHONY: all build clean run test install uninstall

all: clean build

clean:
	$(GO) clean -testcache
	$(RM) -rf bin/*
build:
	$(GO) build -o bin/$(APP) $(LDFLAGS) $(SRC_DIR)/*.go
run:
	$(GO) run $(LDFLAGS) $(SRC_DIR)/*.go
test:
	$(GO) test -v ./...
install: build
	install -d $(PREFIX)/bin
	install -m 755 bin/$(APP) $(PREFIX)/bin/$(APP)
uninstall:
	$(RM) $(PREFIX)/bin/$(APP)
