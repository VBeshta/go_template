build:
	go build -o bin/YOUR-REPO-NAME

run: build
	./bin/YOUR-REPO-NAME

test: 
	go test -v ./...