# vi: ft=make
.PHONY: start test

start:
	go run main.go

test:
	go test -v proxy/*

