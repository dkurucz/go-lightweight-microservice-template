# Go Lightweight Microservice Template

## Overview

This template serves as a lightweight server template designed for eventual enterprise scale.
The service is provided as a Docker service and can also be run standalone.  Out of the box,
the service is coupled with a mysql database and uses GORM to do lightweight queries and commits.
Also out of the box is a simple authentication server using bcrypt to perform text-based encrypted
password comparisons and storage.

## Building

Simply run `make build` to initiate the go build and run the docker build.

## Running

Run `docker compose up` to initiate the server and mock database
Note: if database is not up already, you may need to call compose up a second time after the database is ready in order
to start the API server.

## Testing

This project contains sample Postman collections to invoke and test the APIs directly.  These can be loaded into postman
and executed individually to test the single endpoints.

The current supported endpoints are the following

- GET - "/api/account/"
- GET - "/api/account/{id}"
- POST - "/api/account/create"
- POST - "/api/account/login"
- DELETE - "/api/account/{id}"