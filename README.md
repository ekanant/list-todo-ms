# List To Do Microservice

This is an learning project to implete GRPC with REST

---

## Step
1. Create `tools.go` with this content
    ```go
    // +build tools

    package tools

    import (
        _ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway"
        _ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2"
        _ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
        _ "google.golang.org/protobuf/cmd/protoc-gen-go"
    )
    ```

2. Install protoc https://grpc.io/docs/protoc-installation/
3. Follow step from https://grpc-ecosystem.github.io/grpc-gateway/docs/tutorials/introduction/
---
## Update protobuf file
Run this command
```sh
make generate
```
---
## Reference

https://earthly.dev/blog/golang-grpc-gateway/

https://grpc-ecosystem.github.io/grpc-gateway/docs/tutorials/introduction/