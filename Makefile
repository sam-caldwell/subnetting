clean:
	@rm -rf build &>/dev/null || true
	@mkdir -p build &>/dev/null || true

lint:
	@command -v golint &>/dev/null || go install golang.org/x/lint/golint
	@go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow@latest
	@go vet ./...
	@golint -set_exit_status ./...

test: lint
	go test -v -count=1 ./...

build: clean
	@go build -o build/calculateSubnets ./cmd/calculateSubnets/main.go
