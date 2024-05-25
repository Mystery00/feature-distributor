gen-pb-server:
	protoc --go_opt=Mcommon/protobuf/project.proto=pb/ \
		--go_out=core/ \
		--go-grpc_opt=Mcommon/protobuf/project.proto=pb/ \
		--go-grpc_out=core/ \
		common/protobuf/project.proto