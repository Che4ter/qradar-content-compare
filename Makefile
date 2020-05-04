LDFLAGS=-ldflags "-s -w -X main.Version=$(shell git describe --abbrev=0  --always --tags)"

build:
	go build -o bin/qradar_content_compare_debug main.go

run:
	go run main.go

compile:
	# Linux
	GOOS=linux GOARCH=amd64 go build -o bin/qradar_content_compare ${LDFLAGS} main.go
	# Windows
	GOOS=windows GOARCH=amd64 go build -o bin/qradar_content_compare.exe ${LDFLAGS} main.go

clean:
	rm -rf qradar_compare_report_*
	rm -f bin/*
