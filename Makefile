run:
	go run main.go
build:
	rm -rf ./bin
	go build -o ./bin/brisk
docker: build
	sudo docker build -t brisk .