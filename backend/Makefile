.PHONY: build
build:
	go build .
	mkdir -p bin/
	mv blog bin/

include scripts/makefile/config.mk
include scripts/makefile/db.mk
include scripts/makefile/docker.mk
include scripts/makefile/test.mk

.PHONY: install-tools
install-tools:
	go mod download
	@cat tools/tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %

####################################
# app
####################################
.PHONY: gomodgen
gomodgen:
	GO111MODULE=on go mod init ${FUNCTION}

.PHONY: run
run:
	go run .
