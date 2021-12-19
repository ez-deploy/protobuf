
# generate proto files.
gen: clean
	protoc ./model/*.proto  --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. --proto_path=.
	protoc ./identity/*.proto  --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. --proto_path=.

clean:
	rm -rf ./identity/*.pb.go
	rm -rf ./model/*.pb.go