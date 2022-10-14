all:
	@echo "Action required, for example \"make generate\""
generate:
	$(eval output_folder := "grpc/generated")
	&& cd proto \
	&& protoc -I . \
		--go_out ../${output_folder} --go_opt paths=source_relative \
		--go-grpc_out ../${output_folder} --go-grpc_opt paths=source_relative \
		--grpc-gateway_out ../${output_folder} --grpc-gateway_opt paths=source_relative \
		*.proto
generate_cert:
	cd certs \
	&& cfssl gencert -initca ca-csr.json | cfssljson -bare ca \
	&& cfssl gencert -ca ca.pem -ca-key=ca-key.pem -config ca-config.json \
               -profile=server server-csr.json | cfssljson -bare server