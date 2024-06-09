.PHONY: generate build run doc validate spec clean help

all: clean generate build

validate:
	swagger validate ./api/swagger/swagger.yaml

spec:
	swagger generate spec -o ./api/swagger/swagger-gen.yaml

build: 
	CGO_ENABLED=0 GOOS=linux go build -v -installsuffix cgo ./cmd/task-mate-server
	
run:
	.task-mate-server --port=8080 --host=0.0.0.0 --config=./configs/app.yaml

run-local:
	go run cmd/task-mate-server/main.go --port=8080

doc: validate
	swagger serve ./api/swagger/swagger.yaml --no-open --host=0.0.0.0 --port=8080 --base-path=/

clean:
	rm -rf task-mate-server
	go clean -i .

generate: validate
	swagger -q generate server --exclude-main -A task-mate -f api/swagger/swagger.yaml -s gen/api -m gen/models --principal models.Principal
	go mod tidy