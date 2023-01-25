TAG := $(shell git rev-parse --short HEAD)
DIR := $(shell pwd -L)

SDCLI_VERSION:=v3.0.0
SDCLI:=docker run -ti --mount src="$(DIR)",target="$(DIR)",type="bind" -w "$(DIR)" \
	asecurityteam/sdcli:$(SDCLI_VERSION)

usage.txt: usage.md
	docker run -i pandoc/core:2.14.1 -f markdown -t plain --wrap=auto < $< > $@

dep: usage.txt
	$(SDCLI) go dep

lint:
	$(SDCLI) go lint

test: dep
	$(SDCLI) go test

integration: ;

coverage:
	$(SDCLI) go coverage

doc: ;

build-dev: ;

build: ;

run: ;

deploy-dev: ;

deploy: ;