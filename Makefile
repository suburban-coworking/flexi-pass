# binaries
GOCMD=go
PACKRCMD=packr
DEPCMD=dep

# packr commands
PACKRBUILD=$(PACKRCMD) build
PACKRCLEAN=$(PACKRCMD) clean
PACKRINSTALL=$(PACKRCMD) install

# Go commamds
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get -v -u

# Dep commands
DEPENSURE=$(DEPCMD) ensure

BINARY_NAME=flexi-pass
BINARY_PATH=$(shell which $(BINARY_NAME))

all: install test
build: 
	$(PACKRBUILD) -o $(BINARY_NAME) -v
install: build
	$(PACKRINSTALL)
test:
	$(GOTEST) -v -cover ./...
clean:
	$(PACKRCLEAN)
	$(GOCLEAN)
	rm -fv $(BINARY_NAME)
	rm -fv $(BINARY_PATH)
run: install
	$(BINARY_NAME)
deps:
	$(GOGET) github.com/rubenv/sql-migrate/...
	$(GOGET) github.com/golang/dep/cmd/dep
	$(GOGET) github.com/gobuffalo/packr/...
	$(DEPENSURE)
