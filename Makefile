migration_up: 
	migrate -path migrations/ -database "postgresql://dev:dev@localhost:5102/postgres?sslmode=disable" -verbose up
 
migration_down: 
	migrate -path migrations/ -database "postgresql://dev:dev@localhost:5102/postgres?sslmode=disable" -verbose down

migration_fix: 
	migrate -path migrations/ -database "postgresql://dev:dev@localhost:5102/postgres?sslmode=disable" force VERSION
