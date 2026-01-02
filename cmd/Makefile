SHELL := /bin/bash
INSTALLPATH ?= /usr/local/bin

.PHONY: all
all: install-binaries

install-binaries:
	@for i in `ls | grep -v Makefile` ; do cd ./$$i ; pwd ; CGO_ENABLED=0 go install ; cd .. ; done

install: install-binaries
	@for i in `ls | grep -v Makefile` ; do sudo cp $$GOPATH/bin/$$i $(INSTALLPATH)/ ; sudo chown root:root $(INSTALLPATH)/$$i ; sudo chmod 755 $(INSTALLPATH)/$$i ; done