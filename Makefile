build:
	govendor sync
	go build -i

run:
	go run main.go

fmt:
	go fmt ./...

devdeps:
	go get -u github.com/kardianos/govendor
	govendor init

test:
	go test --tags user -v ./...
