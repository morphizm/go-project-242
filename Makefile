build:
	go build -o bin/hexlet-path-size ./cmd/hexlet-path-size

lint-fix:
	golangci-lint run --fix

lint:
	golangci-lint run
