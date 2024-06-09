gen-pb-health:
	protoc \
		--go_opt=Mcommon/protobuf/health.proto=pb/ \
		--go_out=core/ \
		--go-grpc_opt=Mcommon/protobuf/health.proto=pb/ \
		--go-grpc_out=core/ \
		common/protobuf/health.proto

gen-pb-core:
	protoc \
		--go_opt=Mcommon/protobuf/project.proto=pb/ \
		--go_opt=Mcommon/protobuf/toggle.proto=pb/ \
		--go_opt=Mcommon/protobuf/user.proto=pb/ \
		--go_out=core/ \
		--go-grpc_opt=Mcommon/protobuf/project.proto=pb/ \
		--go-grpc_opt=Mcommon/protobuf/toggle.proto=pb/ \
		--go-grpc_opt=Mcommon/protobuf/user.proto=pb/ \
		--go-grpc_out=core/ \
		common/protobuf/project.proto common/protobuf/toggle.proto common/protobuf/user.proto

gen-pb-endpoint:
	protoc \
		--go_opt=Mcommon/protobuf/toggle.proto=pb/ \
		--go_opt=Mcommon/protobuf/user.proto=pb/ \
		--go_out=endpoint/ \
		--go-grpc_opt=Mcommon/protobuf/toggle.proto=pb/ \
		--go-grpc_opt=Mcommon/protobuf/user.proto=pb/ \
		--go-grpc_out=endpoint/ \
		common/protobuf/toggle.proto common/protobuf/user.proto