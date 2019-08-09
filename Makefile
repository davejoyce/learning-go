PROJECTNAME=$(shell basename "$(PWD)")

# Go related variables.
GOBASE=$(shell pwd)
GOPATH=$(GOBASE)/vendor:$(GOBASE):$(HOME)/go
GOBIN=$(GOBASE)/bin
GOFILES=$(wildcard *.go)

# Redirect error output to a file, so we can show it in development mode.
STDERR=/tmp/.$(PROJECTNAME)-stderr.txt

# PID file will keep the process id of the server
PID=/tmp/.$(PROJECTNAME).pid

# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent

default: compile

build: $(GOFILES)
	@echo " > Building binaries..."
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/$(PROJECTNAME) .

clean:
	@echo " > Cleaning build artifacts..."
	rm -fr bin

compile: clean build
