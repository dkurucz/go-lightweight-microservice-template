NAME = glmt/go-lightweight-microservice-template
VERSION = 0.1
C_NAME = go-lightweight-microservice-template
PORT = 9010

.PHONY: all build test tag_latest release install clean

all: clean build

build:
	cd ./src && env GOOS=linux GOARCH=amd64 go build -o ../bin/main main.go
	docker build -t $(NAME):$(VERSION) -f Dockerfile .

clean:
	rm -f ./cmd/main/main
	docker rm -vf $(C_NAME)

tag_latest:
	docker tag $(NAME):$(VERSION) $(NAME):latest

release: tag_latest
	@if ! docker images $(NAME) | awk '{ print $$2 }' | grep -q -F $(VERSION); then echo "$(NAME) version $(VERSION) is not yet built. Please run 'make build'"; false; fi
	docker push $(NAME):$(VERSION)