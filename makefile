all: build-app up

build-app:
	docker build -t anymind .

up:
	docker-compose up -d

down:
	docker-compose down