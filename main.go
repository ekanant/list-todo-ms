package main

import (
	"context"
	"list-todo-ms/grpc/generated"
	"log"
	"net/http"
	"strings"

	"list-todo-ms/handler"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func main() {
	serverHandler := handler.NewHandler()

	// Create a gRPC server object
	grpcServer := grpc.NewServer()
	// Attach the List Todo service to the server
	generated.RegisterListToDoServer(grpcServer, serverHandler)

	gwmux := runtime.NewServeMux()
	// Register List Todo
	err := generated.RegisterListToDoHandlerServer(context.Background(), gwmux, serverHandler)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	log.Println("Serving gRPC-Gateway on https://0.0.0.0:8090")
	err = http.ListenAndServeTLS(":8090", "./certs/server.pem", "./certs/server-key.pem", grpcHandlerFunc(*grpcServer, gwmux))
	if err != nil {
		log.Fatalln("Failed to start grpc", err)
	}
}

func grpcHandlerFunc(grpcServer grpc.Server, otherHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.HasPrefix(
			r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	})
}
