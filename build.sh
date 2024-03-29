#!/bin/bash

# excute app
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build --ldflags="-s -w" -o deepl_api_win_x86_64.exe
CGO_ENABLED=0 GOOS=windows GOARCH=arm64 go build --ldflags="-s -w" -o deepl_api_win_arm64.exe
CGO_ENABLED=0 GOOS=linux GOARCH=386 go build --ldflags="-s -w" -o deepl_api_linux_x86
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build --ldflags="-s -w" -o deepl_api_linux_amd64
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build --ldflags="-s -w" -o deepl_api_linux_arm64
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build --ldflags="-s -w" -o deepl_api_darwin_amd64
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build --ldflags="-s -w" -o deepl_api_darwin_arm64

# docker app
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build --ldflags="-s -w" -o app_amd64
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build --ldflags="-s -w" -o app_arm64
