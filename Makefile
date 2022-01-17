build-linux:
	GOOS=linux GOARCH=amd64	go build -o ./build/procman ./cmd
run:
	go run ./cmd/main.go
du:
	docker-compose up
dd:
	docker-compose down
bl-du: build-linux du