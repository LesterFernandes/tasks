PROTOC       = protoc
PROTO_DIR    = api/users
PB_DIR       = pb
PROTO_FILES  = $(wildcard $(PROTO_DIR)/*.proto)

.PHONY: proto clean

proto-users:
	$(PROTOC) --proto_path=api/users \
		--go_out=$(PB_DIR) --go_opt=paths=source_relative \
		--go-grpc_out=$(PB_DIR) --go-grpc_opt=paths=source_relative \
		$(PROTO_FILES)

clean:
	rm -f $(PB_DIR)/*.pb.go

