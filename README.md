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