run:
	go run main.go
build:
	rm -rf ./bin
	go build -o ./bin/brisk
docker: build
	sudo docker build -t brisk .
test:
	rm -rf test-results
	mkdir test-results
	go test -v 2>&1 ./... | go-junit-report > ./test-results/report.xml