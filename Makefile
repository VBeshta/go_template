build:
	go build -o bin/go_template //change go_template to {projectName} 

run: build
	./bin/go_template

test: 
	go test -v ./...