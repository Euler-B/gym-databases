include .env
d = $
container = (docker ps -a | grep \queserapp-db | awk '{print $(d)(1)}')
# ip-address = (docker network inspect bridge | grep Gateway | awk 'print $(d)(2)') TODO: integrar luego esta orden dentro del flujo 
container-db-test = (docker ps -a | grep \queserapp-db-test | awk '{print $(d)(1)}')

run:
	go run cmd/main.go

run-dependencies:
	make drop-database
	make database-up
	echo "Esperar algunos minutos mientras se levanta la base de datos"
	sleep 15
	make populate-database

database-up:
	docker run \
	-e POSTGRES_PASSWORD=${POSTGRES_PASSWORD} \
	-e POSTGRES_USER=${POSTGRES_USER} \
	-p 5432:5432 \
	--name queserapp-db \
	-d \
	postgres:latest

populate-database:
	cd  internal/database/ && psql ${POSTGRES_URL} -f create-tables.sql

drop-database:
	docker rm -f $(d)$(container)

stop-database:
	docker stop $(d)$(container)
	

# tests

run-tests:
	make test-environment-up
	echo Espere unos segundo mientras se levanta el la base de datos de prueba
	sleep 15
	make populate-db-test
	make tests
	make -k stop-test-environment


test-environment-up:
	docker run \
	-e POSTGRES_PASSWORD=${POSTGRES_PASSWORD_TEST} \
	-e POSTGRES_USER=${POSTGRES_USER_TEST} \
	-p 5432:5432 \
	--name queserapp-db-test \
	-d \
	postgres:latest

stop-test-environment:
	docker rm -f $(d)$(container-db-test)

populate-db-test:
	cd  internal/databases/migrations && psql ${POSTGRES_URL_TEST} -f create-tables.sql

tests:
	cd cmd/ && go test -v ./...
