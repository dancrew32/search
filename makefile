make:
	vim makefile

deps:
	go get github.com/RediSearch/redisearch-go/redisearch

build:
	go build search.go

run:
	go run search.go

up:
	docker-compose up -d

down:
	docker-compose down
