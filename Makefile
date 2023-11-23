# BASE COMMANDS ################################################################

default: shell

up:
	docker compose up -d

down:
	docker compose down

build:
	docker compose build

rebuild:
	docker compose build --force-rm --no-cache

restart:
	docker compose restart


# EXEC COMMANDS ################################################################

shell:
	docker compose exec api sh
