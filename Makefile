BINARY := golume
VERSION ?= vlatest

.PHONY: build
build:
	go build -o $(BINARY)-$(VERSION)

.PHONY: install
install: build
	install -m 0755 $(BINARY)-$(VERSION) /usr/local/bin/aminal
	rm -f $(BINARY)-$(VERSION)