PROJECT_ROOT?=$(shell pwd)

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o .out/auth-service

.PHONY: build-ci
build-ci: build
	docker build -f Dockerfile -t auth-service:latest .

.PHONY: clean
clean:
	-rm -rf .out/
