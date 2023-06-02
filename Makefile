build:
	@mkdir -p build &>/dev/null || true
	@go build -o build/calculateSubnets ./cmd/calculateSubnets/main.go