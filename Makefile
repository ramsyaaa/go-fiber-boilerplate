IMAGE_NAME := go-fiber-modular-example

.PHONY: build start

build:
	docker build -t $(IMAGE_NAME) .

start:
	docker compose up --build --abort-on-container-exit
