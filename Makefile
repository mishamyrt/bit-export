VERSION = 0.0.1
ENTRYFILE = main.go

BUILD_DIR = build
BINARY_NAME = bit-exporter

GC = go build -trimpath -ldflags="-X 'bit-exporter/cmd.Version=v$(VERSION)' -s -w"

.PHONY: build-static
build-static:
	CGO_ENABLED=0 $(GC) -o "$(BUILD_DIR)/$(BINARY_NAME)_static" $(ENTRYFILE)

.PHONY: build
build:
	$(GC) -o "$(BUILD_DIR)/$(BINARY_NAME)" $(ENTRYFILE)

.PHONY: build-static
run:
	go run $(ENTRYFILE)