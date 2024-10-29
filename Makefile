migrateup:
	migrate -path migration -database "postgresql://rainataputra:root@localhost:5432/inventory_management?sslmode=disable" -verbose up

migratedown:
	migrate -path migration -database "postgresql://rainataputra:root@localhost:5432/inventory_management?sslmode=disable" -verbose down


.PHONY: migrateup migratedown