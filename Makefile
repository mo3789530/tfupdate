BINARY=tfupdate

defualt: build

build: 
	go build -o ${BINARY} .
