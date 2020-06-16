####################################
# test
####################################
test:
	$(eval RICHIGO := $(shell which richgo > /dev/null; echo $$?))
	$(eval CPU_COUNT := $(shell if [ "$(UNAME)" = "darwin" ]; then sysctl -n hw.ncpu; else grep -ce '^processor\s\+:' /proc/cpuinfo; fi))
	@if [ $(RICHIGO) = 0 ]; then richgo test -v ./...; else go test -parallel $(CPU_COUNT) -v ./...; fi

install_richgo:
	# only osx
	brew tap kyoh86/tap
	brew install richgo

install_mockgen:
	go get -u github.com/golang/mock/mockgen

install_goverage:
	go get -u github.com/haya14busa/goverage

gen_mock:
	sh scripts/sh/test/generate_mock.sh

coverage:
	goverage -coverprofile=coverage.out ./app/application/... ./app/domain/... ./app/infrastructure/... ./app/interfaces/...
	go tool cover -html=coverage.out
