EXE			= vault-tools
PACKAGE = e2x.com/vault-tools
BASE 		= $(GOPATH)/src/$(PACKAGE)
VERSION = 0.0.3

.PHONY : all clean fmt test build

all : clean fmt test build

build : test
    @go go get github.com/GeertJohan/go.rice/rice
	@GOOS=darwin GOARCH=amd64 go build  -o $(GOPATH)/bin/$(EXE)-$(VERSION)-osx $(PACKAGE)
	@-rice append -i $(PACKAGE)/project --exec $(GOPATH)/bin/$(EXE)-$(VERSION)-osx
	@GOOS=linux GOARCH=amd64 go build  -o $(GOPATH)/bin/$(EXE)-$(VERSION)-linux $(PACKAGE)
	@-rice append -i $(PACKAGE)/project --exec $(GOPATH)/bin/$(EXE)-$(VERSION)-linux

test : fmt
	@go test -v -cover ./...

test-junit : fmt
	@go get -u github.com/jstemmer/go-junit-report
	@go test -v 2>&1 ./... | go-junit-report > report.xml
fmt :
	@gofmt -w $(BASE)/*.go

clean :
	@-rm $(GOPATH)/bin/$(EXE)-*
	@-rm $(GOPATH)/bin/rice
