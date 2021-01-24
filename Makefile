.PHONEY: up
## up: docker-compose up
up:
	docker-compose up -d --build

.PHONEY: run
## run: go run
run:
	docker-compose exec app go run main.go

.PHONEY: down
# down: docker-compose down
down:
	docker-compose down

.PHONEY: ps
# ps: docker-compose ps
ps:
	docker-compose ps
