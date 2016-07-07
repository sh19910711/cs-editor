SWAGGER_VERSION=2.1.4

build:
	govendor sync
	go build -i

swagger-ui: tmp/swagger-ui.tar.gz
	cd tmp && tar zxf swagger-ui.tar.gz
	mkdir -p assets/swagger-ui
	cp -r tmp/swagger-ui-$(SWAGGER_VERSION)/dist/* assets/swagger-ui/

tmp/swagger-ui.tar.gz:
	mkdir -p tmp
	cd tmp && curl -L https://github.com/swagger-api/swagger-ui/archive/v$(SWAGGER_VERSION).tar.gz > swagger-ui.tar.gz

run:
	go run main.go

fmt:
	go fmt ./...

devdeps:
	go get -u github.com/kardianos/govendor
	govendor init

test:
	go test --tags user -v ./...
