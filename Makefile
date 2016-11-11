SHELL=/bin/bash
SRCS=$(shell cd src; ls | sed -e 's@github.com@@')
GOPATH=$(shell pwd)
export GOPATH

all: deps
	@for SRC in $(SRCS); do \
		echo "Building $$SRC"; \
	  go install $$SRC; \
	done

deps:
	[ -d src/github.com/fsnotify/fsevents ] || go get -f -u -v github.com/fsnotify/fsevents
	[ -d src/github.com/DHowett/go-plist ]  || go get -f -u -v github.com/DHowett/go-plist

clean:
	rm -rf bin pkg src/github.com

# test runs the test suite and vets the code
#test:
#	go list $(TEST) | xargs -n1 go test -timeout=60s -parallel=10 $(TESTARGS)
