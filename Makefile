PACKAGE  = wifimeidod
GO = go
GOPATH   = $(CURDIR)/.gopath
BASE     = $(GOPATH)/src/$(PACKAGE)

$(BASE):
	@mkdir -p $(dir $@)
	@ln -sf $(CURDIR) $@

.PHONY: all
all: | $(BASE)
	 GOPATH=$(GOPATH) $(GO) build -o bin/$(PACKAGE) wifimeidod

