generate_grpc_logic_stubs_client:
	protoc --proto_path=.. \
  --go_out=../invoicer --go_opt=paths=source_relative \
  --go-grpc_out=../invoicer --go-grpc_opt=paths=source_relative \
  ../invoicer.proto
