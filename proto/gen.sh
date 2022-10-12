work_dir=$(dirname "$0")
cd ${work_dir}
output_folder="grpc/generated"
protoc -I . \
    --go_out ../${output_folder} --go_opt paths=source_relative \
    --go-grpc_out ../${output_folder} --go-grpc_opt paths=source_relative \
    --grpc-gateway_out ../${output_folder} --grpc-gateway_opt paths=source_relative \
    *.proto
echo "generate complete"