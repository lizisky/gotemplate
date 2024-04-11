.PHONY: build
build:
	go build -ldflags="-s -w" -o bin/liziskyd src/main.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t lizitime/liziskyd_experience:latest

docker_pro:
	docker build . -t lizitime/liziskyd:latest

.PHONY: clean
clean:
	rm bin/liziskyd
	rm -rf bin/logs
