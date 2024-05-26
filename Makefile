gen-pb-server:
	protoc \
		--go_opt=Mcommon/protobuf/project.proto=pb/ \
		--go_opt=Mcommon/protobuf/toggle.proto=pb/ \
		--go_out=core/ \
		--go-grpc_opt=Mcommon/protobuf/project.proto=pb/ \
		--go-grpc_opt=Mcommon/protobuf/toggle.proto=pb/ \
		--go-grpc_out=core/ \
		common/protobuf/project.proto common/protobuf/toggle.proto