gen-pb-server:
	protoc --go_opt=Mprotobuf/project.proto=pb/ \
		--go_out=core/ \
		--go-grpc_opt=Mprotobuf/project.proto=pb/ \
		--go-grpc_out=core/ \
		protobuf/project.proto