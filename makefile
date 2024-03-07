include internal/database/.env
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
	make db-migration

database-up:
	docker run \
	-e POSTGRES_PASSWORD=${POSTGRES_PASSWORD} \
	-e POSTGRES_USER=${POSTGRES_USER} \
	-p 5432:5432 \
	--name queserapp-db \
	-d \
	postgres:latest

db-migration:
	migrate -path internal/database/migrations/ -database ${POSTGRES_URL} up

db-rollback:
	migrate -path internal/database/migrations/ -database ${POSTGRES_URL} down

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
	-p 14256:5432 \
	--name queserapp-db-test \
	-d \
	postgres:latest

stop-test-environment:
	docker rm -f $(d)$(container-db-test)

populate-db-test:
	migrate -path internal/database/migrations/ -database ${POSTGRES_URL_TEST} up

tests:
	go test -v ./...
