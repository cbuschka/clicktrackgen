.PHONY:	build run
build:
	mkdir -p dist/
	go build -o dist/clicktrackgen cmd/clicktrackgen.go

run:
	go run cmd/clicktrackgen.go
