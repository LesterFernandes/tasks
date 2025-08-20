PROTOC       = protoc
PROTO_DIR    = api/users
PB_DIR       = pb

PROTO_FILES  = $(wildcard $(PROTO_DIR)/*.proto)

.PHONY: proto clean run

proto:
	$(PROTOC) --proto_path=$(PROTO_DIR) \
		--go_out=$(PB_DIR) --go_opt=paths=source_relative \
		--go-grpc_out=$(PB_DIR) --go-grpc_opt=paths=source_relative \
		$(PROTO_FILES)

clean:
	rm -f $(PB_DIR)/*.pb.go

run:
	go run ./cmd/server
