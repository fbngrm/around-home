#!/usr/bin/env just --justfile

db_user := 'postgres'
db_name := 'match'

default:
	just -l

build: generate
	cd server && go build
	@echo "server built"
	cd cmd/client && go build
	@echo "client built"
	cd cmd/migrate && go build
	@echo "migrate built"
	@just lint

lint:
	golangci-lint run

release:
	just buf breaking --against ".git#branch=main,subdir=."

setup: setup-go generate

generate:
	make gen/proto/go/match/v1/*.go >/dev/null
	make ent/ent.go >/dev/null

seed-and-run:
    docker-compose --profile seed up

run:
    docker-compose up

run-server:
	cd server && go build
	DB_NAME={{db_name}} DB_USER={{db_user}} ./server/server

stop:
    docker-compose down

client: generate
	cd cmd/client && go build
	./cmd/client/client

seed: generate
	cd cmd/seed && go build
	DB_NAME={{db_name}} DB_USER={{db_user}} ./cmd/seed/seed

migrate *args='': generate
	cd cmd/migrate && go build
	DB_NAME={{db_name}} DB_USER={{db_user}} ./cmd/migrate/migrate {{args}}

buf *args='':
	cd apis && PATH="$PATH:$(go env GOPATH)/bin" buf {{args}}

ent *args='':
	go run entgo.io/ent/cmd/ent {{args}}

setup-go:
	go install entgo.io/ent/cmd/ent@latest
	cd gen && go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	cd gen && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	cd gen && go install \
	  github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
	  github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
	cd gen && go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
