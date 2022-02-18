
mkdir -p pkg/protogen
protoc --go_out=./pkg/protogen \
    --go-grpc_out=./pkg/protogen \
    protos/login.proto