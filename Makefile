# Variables
DOCKER_COMPOSE_FILE := docker-compose.yaml
DOCKER_IMAGE_NAME := asif/travel-booking-portal
DOCKER_IMAGE_TAG := 1.0

# Targets
build:
	docker build -t $(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG) .

run:
	@if docker image inspect $(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG) > /dev/null 2>&1; \
	then \
		echo "Docker image exists. Starting containers..."; \
	else \
		echo "Docker image not found. Building it..."; \
		make build; \
	fi
	docker-compose -f $(DOCKER_COMPOSE_FILE) up

stop:
	docker-compose -f $(DOCKER_COMPOSE_FILE) down

test: run
	docker-compose -f $(DOCKER_COMPOSE_FILE) exec -e EMAIL_SMTP_USERNAME=testuser -e EMAIL_SMTP_PASSWORD=testpassword api go test ./...

format:
	bash ./scripts/format.sh

check: format
	bash ./scripts/check.sh

db_prepare:
	docker-compose -f $(DOCKER_COMPOSE_FILE) exec -T mongodb mongosh --eval "use travel-booking-portal"

.PHONY: build run stop test format check db_prepare
