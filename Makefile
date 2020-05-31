GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOCLEAN=$(GOCMD) clean
BINARY_DIR=$(PWD)/cmd/graceful-restarter
EXAMPLE_DIR=$(PWD)/example
LISTENER_DIR=$(PWD)/graceful-listener
MAKEFLAGS += --no-builtin-rules
.SUFFIXES:

.PHONY: all
all: test build

build: sample grestarter

test: listener_test

grestarter: $(BINARY_DIR)/main.go
	$(GOBUILD) -o $(BINARY_DIR)/$@ -v $<

sample: $(EXAMPLE_DIR)/example.go
	$(GOBUILD) -o $(EXAMPLE_DIR)/$@ -v $<

.PHONY: clean
clean:
	$(RM) $(BINARY_DIR)/grestarter
	$(RM) $(EXAMPLE_DIR)/sample

listener_test: $(LISTENER_DIR)/listener_test.go
	$(GOTEST) -v $(LISTENER_DIR)
