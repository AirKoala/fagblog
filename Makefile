# Project name
PROJECT_NAME=fagblog

# Can be overriden by passing in GOOS and GOARCH for cross compile
GOOS ?= linux
GOARCH ?= amd64
BINARY_NAME := $(PROJECT_NAME)-$(GOOS)-$(GOARCH)

ifdef GOARM
	BINARY_NAME := $(BINARY_NAME)v$(GOARM)
endif

# Install paths
PREFIX := /usr/local
BINDIR := $(PREFIX)/bin
DATADIR := $(PREFIX)/share/$(PROJECT_NAME)
CONTENTDIR := /var/lib/$(PROJECT_NAME)
CONFIGDIR := /etc/$(PROJECT_NAME)
SERVICE_PATH := /etc/systemd/system/$(PROJECT_NAME).service

PACKAGE_INCLUDE := $(shell find static/) $(shell find templates/) Makefile $(BINARY_NAME) $(PROJECT_NAME).service LICENSE default_config.toml
SRC_FILES := $(shell find . -name '*.go')

.PHONY: run build install uninstall clean package

## Run the project
run:
	go run .

## Build the project
build: $(BINARY_NAME)

$(BINARY_NAME): $(SRC_FILES)
ifeq ($(strip $(SRC_FILES)),)
	@echo No source files found. Skipping.
else
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(BINARY_NAME) .
endif

## Clean build files
clean:
	rm -f $(PROJECT_NAME)-*

## Install the binary and resource files
install: $(PACKAGE_INCLUDE)
	# Create user and group if they don't exist
	@if ! id -u $(PROJECT_NAME) >/dev/null 2>&1; then \
		echo "Creating user: $(PROJECT_NAME)"; \
		useradd --system --no-create-home --shell /usr/sbin/nologin $(PROJECT_NAME); \
	fi

	install -Dm755 -s $(BINARY_NAME) $(BINDIR)/$(PROJECT_NAME)

	install -d -m 755 -o root -g $(PROJECT_NAME) $(DATADIR)
	cp -r static templates $(DATADIR)
	chown -R root:$(PROJECT_NAME) $(DATADIR)
	chmod -R 644 $(DATADIR)

	# Set all directories to be searchable.
	chmod ug+x $(find $DATADIR -type d)

	install -d -m 750 -o root -g $(PROJECT_NAME) $(CONFIGDIR)
	install -m 640 -o root -g $(PROJECT_NAME) default_config.toml $(CONFIGDIR)/config.toml

	install -d -m 755 -o root -g $(PROJECT_NAME) $(CONTENTDIR)

	# Install systemd service
	install -m 644 $(PROJECT_NAME).service $(SERVICE_PATH)
	systemctl daemon-reload

## Uninstall the binary and resource files
uninstall:
	@echo "This will remove:"
	@echo "  - $(BINDIR)/$(PROJECT_NAME)"
	@echo "  - $(DATADIR)/"
	@echo "  - $(SERVICE_PATH)"
	@read -p "Are you sure you want to uninstall $(PROJECT_NAME)? [y/N] " confirm; \
	if [ "$$confirm" = "y" ] || [ "$$confirm" = "Y" ]; then \
		echo "Uninstalling..."; \
		rm -f $(BINDIR)/$(PROJECT_NAME); \
		rm -rf $(DATADIR) $(CONFIGDIR); \
		rm -f $(SERVICE_PATH); \
		systemctl daemon-reload; \
		echo "Done."; \
	else \
		echo "Uninstall cancelled."; \
	fi

## Create an xzipped tarball
$(BINARY_NAME).tar.xz: $(PACKAGE_INCLUDE)
	# Create a tarball with the binary and resource files
	tar -cJf $(BINARY_NAME).tar.xz $(PACKAGE_INCLUDE)

package: $(BINARY_NAME).tar.xz
