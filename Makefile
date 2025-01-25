.PHONY: server client

server:
	protoc --go_out=./server/internal --proto_path=. api.proto

client:
	protoc \
		--plugin=protoc-gen-ts_proto=$(shell which protoc-gen-ts_proto) \
		--ts_proto_out=./client/src/api \
		--proto_path=. \
		api.proto


proto: server client

conn:
	psql -h localhost -p 5432 -U admin -d mydb

up:
	docker compose up -d
down:
	docker compose down
clean:
	docker compose down --volumes --remove-orphans