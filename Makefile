# ==============================================================================
# Docker support for server

server-build:
	@docker build -t onecv_server .

server-run:
	@docker run -p 8080:8080 onecv_server

# ==============================================================================
# Docker support for db

db:
	@docker run -t -p 5431:5432 --env-file=.env postgres:14

# ==============================================================================
# Docker-compose support

compose-build:
	@docker compose build --no-cache

compose-up:
	@docker compose up --quiet-pull --remove-orphans

compose-down:
	@docker compose down --remove-orphans

start: compose-build compose-up

restart: compose-down compose-build compose-up