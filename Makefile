# Makefile for Terminal Chat

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOTEST=$(GOCMD) test
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get
BINARY_NAME=myapp
BINARY_UNIX=$(BINARY_NAME)_unix

# All target
all: test build

# Build the project
build:
	$(GOBUILD) -o $(BINARY_NAME) -v

# Run the project
run:
	$(GORUN) .

# Test the project
test:
	$(GOTEST) -v ./...

# Clean the project
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

# Install the project
install:
	$(GOBUILD) -o $(BINARY_UNIX) -v
	sudo mv $(BINARY_UNIX) /usr/local/bin/$(BINARY_NAME)

# Download dependencies
deps:
	$(GOGET) -u ./...

.PHONY: all build run test clean install deps
