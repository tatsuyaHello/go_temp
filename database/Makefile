export GO111MODULE := off

DBNAME:=tatsuya_db
DB_HOST:=127.0.0.1
DB_PORT:=3306
DATASOURCE:=root:password@tcp($(DB_HOST):$(DB_PORT))/$(DBNAME)

install:
	which goose || GO111MODULE=off go get -u github.com/pressly/goose/cmd/goose

docker-compose/up:
	docker-compose up -d

docker-compose/down:
	docker-compose down

mysql:
	mysql -u root -h localhost --protocol tcp -p $(DBNAME)

migrate/init:
	mysql -u root -h localhost --protocol tcp -e "create database \`$(DBNAME)\`" -p

migrate/status:
	@goose -dir migrations mysql "$(DATASOURCE)" status

migrate/up:
	@goose -dir migrations mysql "$(DATASOURCE)" up

__migrate/down:
	@goose -dir migrations mysql "$(DATASOURCE)" down

init: install docker-compose/up migrate/init migrate/up
