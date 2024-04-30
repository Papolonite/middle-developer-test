MIGRATION_PATH = $(PWD)/platform/migrations
DATABASE_URL = postgresql://dev:dev@localhost:5102/postgres?sslmode=disable

migrate.up: 
	migrate -path ${MIGRATION_PATH} -database "${DATABASE_URL}" -verbose up
 
migrate.down: 
	migrate -path ${MIGRATION_PATH} -database "${DATABASE_URL}" -verbose down

migrate.fix: 
	migrate -path ${MIGRATION_PATH} -database "${DATABASE_URL}" force ${VERSION}