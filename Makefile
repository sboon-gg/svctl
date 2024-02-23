generate:
	go generate ./...

build: build-linux build-win

build-linux: generate
	GOOS=linux CGO_ENABLED=0 go build -o bin/svctl ./

build-win: generate
	GOOS=windows CGO_ENABLED=0 go build -o bin/svctl.exe ./

test: generate
	go run gotest.tools/gotestsum@v1.11.0 -- -count=1 ./...

lint:
	go run github.com/golangci/golangci-lint/cmd/golangci-lint run
