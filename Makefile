#!make

# DOCKER
docker-up: 
	docker-compose up -d

docker-down: 
	docker-compose down

docker-build:
	docker build --tag dealls .

# MIGRATION
migrate: export ENVIRONTMENT=$(env)
migrate:
	@bash ./.bash/migrate.sh


# DEVELOPMENT
serve: export ENVIRONTMENT=$(env)
serve:
	@bash ./.bash/serve.sh

# TESTING
tester: export ENVIRONTMENT=$(env)
tester:
	@bash ./.bash/test.sh