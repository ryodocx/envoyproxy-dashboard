#!/usr/bin/make -f

build: go-build

go-build: .tmp/bin/*

.tmp/bin/*: .tmp/dist/index.html go.sum backend/db/client.go
	go build -o .tmp/bin/ .

go.sum:
	go mod tidy

backend/db/client.go: backend/db/schema/*.go
	go generate ./...

.tmp/dist/index.html: node_modules/.package-lock.json
	npm run build

node_modules/.package-lock.json:
	npm install
