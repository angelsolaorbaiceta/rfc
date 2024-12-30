GO_FILES = rfc.go
BIN_NAME = rfc

.PHONY: install

all: install

# Install the program
install: build
	go install

# Build the program
build: $(GO_FILES)
	go build -o $(BIN_NAME)


