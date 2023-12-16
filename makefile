BINARY_NAME=bible_reader

run:
	go run main.go

build:
	go build -o ${BINARY_NAME} ./src/main.go

clean:
	go clean
	rm ${BINARY_NAME}
