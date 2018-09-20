
# Go parameters
GOCMD=go
DEPCMD=dep
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOINSTALL=$(GOCMD) install
DEPENSURE=$(DEPCMD) ensure
BINARY_NAME=flexi-pass

all: test build
build: 
	$(GOBUILD) -o $(BINARY_NAME) -v
install: build
	$(GOINSTALL) $(BINARY_NAME)
test: build
	$(GOTEST) -v ./...
clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
run: install
	$(BINARY_NAME)
deps:
	$(DEPENSURE)
