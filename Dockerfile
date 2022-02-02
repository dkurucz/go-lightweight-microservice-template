FROM golang:1.17
WORKDIR /var/opt/go-lightweight-microservice-template/

# first set up the OS
RUN \
    apt-get update \
    && apt-get -y install vim \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

EXPOSE 9010

# layer in executable
COPY ./cmd/main/main /glmt

# optional: copy sources and build directly into the container
#COPY src src
#RUN cd cmd/main && go get -d -v ./...
#RUN cd cmd/main && go install -v ./...
#RUN cd cmd/main && go build -o /glmt

CMD [ "/glmt" ]
