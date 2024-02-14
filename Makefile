generate:
	go generate ./...

build: build-linux build-win

build-linux: generate
	GOOS=linux CGO_ENABLED=0 go build -o bin/svctl ./

build-win: generate
	GOOS=windows CGO_ENABLED=0 go build -o bin/svctl.exe ./
