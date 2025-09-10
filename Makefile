
test_short:
	@echo '-------------------'
	@echo 'RUN GO TEST SUMMARY'
	@echo '-------------------'
	@go clean -testcache
	@go test -v ./... | grep -E '^(ok|FAIL)'
	@echo '-------------------'

test:
	go clean -testcache
	go test -v ./...

run:
	go run ./cmd/gochess.go

.PHONY: build
build:
	rm -rf ./build/
	mkdir -p ./build/
	go build \
		-o ./build/gochess \
		-gcflags -m=2 \
		./cmd/ 

install:
	ln "$(realpath ./build/gochess)" -s ~/.local/bin/gochess

hub_update:
	@hub_ctrl ${HUB_MODE} ln "$(realpath ./build/gochess)"

