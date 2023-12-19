BINARY_NAME=bible_reader

run:
	templ generate
	go run cmd/main.go

templ:
	templ generate

build:
	go build -o ${BINARY_NAME} ./src/main.go

clean:
	go clean
	rm ${BINARY_NAME}
