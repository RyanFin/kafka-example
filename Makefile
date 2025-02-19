run:
	go run main.go

docker:
	docker-compose up -d

producer:
	go run pkg/producer/producer.go

consumer:
	go run pkg/consumer/consumer.go

.PHONY: run docker producer consumer