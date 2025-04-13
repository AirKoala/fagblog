# Project name
PROJECT_NAME=fagblog

# Can be overriden by passing in GOOS and GOARCH for cross compile
GOOS ?= linux
GOARCH ?= amd64
BINARY_NAME=$(PROJECT_NAME)-$(GOOS)-$(GOARCH)

# Install paths
PREFIX=/usr/local
BINDIR=$(PREFIX)/bin
DATADIR=$(PREFIX)/share/$(PROJECT_NAME)

.PHONY: run build build-cross install clean

## Run the project
run:
	go run .

## Build the project
build: $(BINARY_NAME)

$(BINARY_NAME):
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(BINARY_NAME) .

## Clean build files
clean:
	rm -f $(PROJECT_NAME)-*

## Install the binary and resource files
install: $(BINARY_NAME)
	install -Dm755 $(BINARY_NAME) $(BINDIR)/$(PROJECT_NAME)
	install -d $(DATADIR)/static
	install -d $(DATADIR)/templates
	cp -r static/* $(DATADIR)/static/
	cp -r templates/* $(DATADIR)/templates/

## Uninstall the binary and resource files
uninstall:
	@echo "This will remove:"
	@echo "  - $(BINDIR)/$(PROJECT_NAME)"
	@echo "  - $(DATADIR)/"
	@read -p "Are you sure you want to uninstall $(PROJECT_NAME)? [y/N] " confirm; \
	if [ "$$confirm" = "y" ] || [ "$$confirm" = "Y" ]; then \
		echo "Uninstalling..."; \
		rm -f $(BINDIR)/$(PROJECT_NAME); \
		rm -rf $(DATADIR); \
		echo "Done."; \
	else \
		echo "Uninstall cancelled."; \
	fi

## Create an xzipped tarball
package: build
	# Create a tarball with the binary and resource files
	tar -cJvf $(BINARY_NAME).tar.xz $(BINARY_NAME) static/ templates/ Makefile
