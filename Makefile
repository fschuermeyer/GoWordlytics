build:
	go build -o bin/gowordlytics cmd/root.go

test:
	go test -v ./...

install:
	go install
