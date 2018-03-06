#! /usr/bin/make
#

CURRENT_DIR := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

all: clean generate build

dep:
	@dep ensure -update

clean:
	@rm -rf app
	@rm -rf client
	@rm -rf tool
	@rm -rf swagger
	@rm -f recipe

generate:
	# @rm -rf vendor/github.com/goadesign
	@goagen bootstrap -d github.com/jaredwarren/recipe/design

build:
	@CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -o recipe -ldflags '-w' .

ae-build:
	@if [ ! -d $(HOME)/recipe ]; then \
		mkdir $(HOME)/recipe; \
		ln -s $(CURRENT_DIR)/appengine.go $(HOME)/recipe/appengine.go; \
		ln -s $(CURRENT_DIR)/app.yaml     $(HOME)/recipe/app.yaml; \
	fi

ae-deploy: ae-build
	cd $(HOME)/recipe
	gcloud app deploy --project recipe