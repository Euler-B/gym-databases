include .env
d = $
container = (docker ps -a | grep \tutorial | awk '{print $(d)(1)}')

run:
	go run main.go

database-up:
	make drop-database
	docker compose up -d 
	echo "Espere algunos minutos mientras se levanta el contenedor de la base de datos"

drop-database:
	docker rm $(d)$(container)

stop-database:
	docker stop $(d)$(container)

tests:
	go test -v ./...
