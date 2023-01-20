version = v1.0
repo = quay.io/redhattraining
local_address = $(shell hostname -i | cut -f 2 -d ' ')

all: fmt build run

fmt:
	gofmt -w *.go

build:
	buildah bud -t exoplanets:latest .

run:
	podman run \
		-p 8080:8080 \
		-e DB_HOST=$(local_address) \
		-e DB_PORT=5432 \
		-e DB_USER=user \
		-e DB_PASSWORD=password \
		-e DB_NAME=postgres \
		exoplanets:latest

run-without-db:
	podman run \
		-p 8080:8080 \
		exoplanets:latest

tag:
	buildah tag exoplanets:latest $(repo)/exoplanets:latest
	buildah tag exoplanets:latest $(repo)/exoplanets:$(version)

push:
	podman push $(repo)/exoplanets:latest
	podman push $(repo)/exoplanets:$(version)

pg-up:
	podman run -d \
		--name postgres \
		-p $(local_address):5432:5432 \
		-e POSTGRES_PASSWORD=password \
		-e POSTGRES_USER=user \
		postgres:12-alpine

pg-down:
	podman rm -f postgres

.PHONY: fmt build run run-without-db tag push pg-up pg-down
