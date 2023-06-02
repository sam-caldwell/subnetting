clean:
	@rm -rf build &>/dev/null || true
	@mkdir -p build &>/dev/null || true

build: clean
	@go build -o build/calculateSubnets ./cmd/calculateSubnets/main.go