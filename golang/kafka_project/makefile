# ---- Configuration ----
TOPIC=my-topic
BROKER=localhost:9092

# ---- Docker Compose ----
up:
	docker-compose up -d

down:
	docker-compose down

logs:
	docker-compose logs -f

# ---- Kafka Topics ----
create-topic:
	docker exec -it kafka kafka-topics --create \
		--topic $(TOPIC) \
		--bootstrap-server $(BROKER) \
		--partitions 1 \
		--replication-factor 1

list-topics:
	docker exec -it kafka kafka-topics --list \
		--bootstrap-server $(BROKER)

describe-topic:
	docker exec -it kafka kafka-topics --describe \
		--topic $(TOPIC) \
		--bootstrap-server $(BROKER)

delete-topic:
	docker exec -it kafka kafka-topics --delete \
		--topic $(TOPIC) \
		--bootstrap-server $(BROKER)

describe-group:
	docker exec -it kafka kafka-consumer-groups \
		--bootstrap-server $(BROKER) \
		--describe \
		--group $(GROUP)

list-groups:
	docker exec -it kafka kafka-consumer-groups \
		--bootstrap-server $(BROKER) \
		--list

# ---- Go App ----
run:
	go run ./main.go

tidy:
	go mod tidy

fmt:
	go fmt ./...

lint:
	golangci-lint run

test:
	go test ./...
