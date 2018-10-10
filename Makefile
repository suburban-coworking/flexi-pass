
# Go parameters
GOCMD=go
PACKRCMD=packr
DEPCMD=dep
GOBUILD=$(PACKRCMD) build
GOCLEAN=$(GOCMD) clean
PACKRCLEAN=$(PACKRCMD) clean
GOTEST=$(GOCMD) test
GOINSTALL=$(PACKRCMD) install
DEPENSURE=$(DEPCMD) ensure
BINARY_NAME=flexi-pass
BINARY_PATH=$(shell which $(BINARY_NAME))

all: install test
build: 
	$(GOBUILD) -o $(BINARY_NAME) -v
install: build
	$(GOINSTALL)
test:
	$(GOTEST) -v ./...
clean:
	$(PACKRCLEAN)
	$(GOCLEAN)
	rm -fv $(BINARY_NAME)
	rm -fv $(BINARY_PATH)
run: install
	$(BINARY_NAME)
deps:
	$(DEPENSURE)
