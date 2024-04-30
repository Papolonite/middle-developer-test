migration_up: 
	migrate -path platform/migrations/ -database "postgresql://dev:dev@localhost:5102/postgres?sslmode=disable" -verbose up
 
migration_down: 
	migrate -path platform/migrations/ -database "postgresql://dev:dev@localhost:5102/postgres?sslmode=disable" -verbose down

migration_fix: 
	migrate -path platform/migrations/ -database "postgresql://dev:dev@localhost:5102/postgres?sslmode=disable" force VERSION
