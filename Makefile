# vi: ft=make
.PHONY: start test

start:
	go run main.go

test:
	go test -v proxy/*

local_test:
	bin/start_chromium.sh
	@make test
	bin/stop_chromium.sh
