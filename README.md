
# Chrome Headless Proxy

[![Build Status](https://travis-ci.org/cberez/chrome-headless-proxy.svg?branch=master)](https://travis-ci.org/cberez/chrome-headless-proxy)

Simple Go proxy to secure access to one or more chrome headless instancess. If multiple chrome headless instances are specified, it will randomly choose one.

## Build

```bash
go build
```

## Run

```bash
make start
# or
go run -addresses=host:port,... -port=8080 -timeout=5 -apikey="some api key"
```

`-addresses` expects chrome headless addresses.

## Test

```bash
make test # expects a chrome headless running on localhost:9222
# or
make local_test # expects chromium installed
```

