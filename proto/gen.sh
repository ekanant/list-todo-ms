work_dir=$(dirname "$0")
cd ${work_dir}
protoc -I . \
    --go_out ../handler --go_opt paths=source_relative \
    --go-grpc_out ../handler --go-grpc_opt paths=source_relative \
    --grpc-gateway_out ../handler --grpc-gateway_opt paths=source_relative \
    *.proto
echo "generate complete"