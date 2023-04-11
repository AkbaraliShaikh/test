all:
	rm -rf parser
	go generate ./...
	go run main.go

clean:
	rm -rf parser

generate: clean
	go generate ./...

build:
	go build

run: build
	./rule-engine

test: build
	go test ./... -v

bench: build
	go test -run=Bench -benchmem -bench=.

benchm: build
	go test -run=Bench -benchmem -bench=. -benchtime=60s