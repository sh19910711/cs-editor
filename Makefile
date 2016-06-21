build:
	go build -i

run:
	go run *.go

devdeps:
	go get -u github.com/kardianos/govendor
	govendor init

test:
	go test -v ./...
