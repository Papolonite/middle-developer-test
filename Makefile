MIGRATION_PATH = platform/migrations/
DATABASE_URL = postgresql://dev:dev@localhost:5102/postgres?sslmode=disable

migrate.up: 
	migrate -path ${MIGRATION_PATH} -database "${DATABASE_URL}" -verbose up
 
migrate.down: 
	migrate -path ${MIGRATION_PATH} -database "${DATABASE_URL}" -verbose down

migrate.fix: 
	migrate -path ${MIGRATION_PATH} -database "${DATABASE_URL}" force ${VERSION}

app.test:
	go test -v ./..

app.run:
	go run main.go

docker.run:
	docker compose up --build

docker.stop :
	docker compose stop
	docker compose down

docker.test:
	docker-compose -f docker-compose.test.yml up --build api --abort-on-container-exit
	docker compose -f docker-compose.test.yml down