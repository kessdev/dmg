GO_CMD=go
GO_BUILD=$(GO_CMD) build
GO_CLEAN=$(GO_CMD) clean
GO_TEST=$(GO_CMD) test

ANTLR_CMD=java -jar ./java-libs/antlr-4.11.1-complete.jar

BINARY_FOLDER=bin

compile: test build
build:
	$(GO_BUILD) -o $(BINARY_FOLDER) -v
test:
	$(GO_TEST) -v ./...
generate:
	$(ANTLR_CMD) -Dlanguage=Go -no-visitor -package parser -o parser -Xexact-output-dir ./antlr-grammars/*.g4