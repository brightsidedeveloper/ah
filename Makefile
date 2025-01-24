.PHONY: server client

server:
	protoc --go_out=./server/internal --proto_path=. api.proto

client:
	protoc \
		--plugin=protoc-gen-ts_proto=$(shell which protoc-gen-ts_proto) \
		--ts_proto_out=./client/src/api \
		--proto_path=. \
		api.proto


all: server client