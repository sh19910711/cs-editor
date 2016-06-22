build:
	go build -i

run:
	govendor sync
	go run *.go

devdeps:
	go get -u github.com/kardianos/govendor
	govendor init

test:
	go test -v ./...
