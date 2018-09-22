
# Go parameters
GOCMD=go
DEPCMD=dep
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOINSTALL=$(GOCMD) install
DEPENSURE=$(DEPCMD) ensure
BINARY_NAME=flexi-pass
BINARY_PATH=$(shell which $(BINARY_NAME))

all: install test
build: 
	$(GOBUILD) -o $(BINARY_NAME) -v
install: build
	$(GOINSTALL) -v
test:
	$(GOTEST) -v ./...
clean:
	$(GOCLEAN)
	rm -fv $(BINARY_NAME)
	rm -fv $(BINARY_PATH)
run: install
	$(BINARY_NAME)
deps:
	$(DEPENSURE)
