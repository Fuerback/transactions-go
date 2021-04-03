img := transaction-go

docker-image:
	docker build -t $(img)

run-local: 
	go run .

docker-run:
	docker run -p 8080:8000 $(img)

build:
	go build .

unit-test:
	go test .\service\... .\repository\...

integration-test:
	go test .\controller\...