make:
	vim makefile

deps:
	go get github.com/RediSearch/redisearch-go/redisearch && \
	go get github.com/gorilla/mux
	go get github.com/stretchr/testify/assert

build:
	go build search.go

run:
	go run search.go

up:
	docker-compose up -d

down:
	docker-compose down

test:
	go test ./...

format:
	gofmt -w .

checkin: format test
	git add -Ap && git commit && git push
