run:
	go run cmd/main.go

proto-gen:
	
	./scripts/gen-proto.sh

tidy:
	go mod tidy
	go mod vendor

migrate_up:
	migrate -path migrations/ -database postgres://ahrorbek:3108@localhost:5432/user_db up


migrate_down:
	migrate -path migrations/ -database postgres://ahrorbek:3108@localhost:5432/user_db down


migrate_force:
	migrate -path migrations/ -database postgres://ahrorbek:3108@localhost:5432/user_db force 1