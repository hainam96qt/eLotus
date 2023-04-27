protoc --proto_path=api --go_out=api --go_opt=paths=source_relative --go-grpc_out=api --go-grpc_opt=paths=source_relative service.proto && protoc -I . --proto_path=api --grpc-gateway_out .\
    --grpc-gateway_opt logtostderr=true \
    --grpc-gateway_opt paths=source_relative \
    --grpc-gateway_opt grpc_api_configuration="api/config/config.yaml" \
    --grpc-gateway_opt standalone=true \
    api/service.proto
