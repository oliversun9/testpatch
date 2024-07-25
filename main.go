package main

// import "buf.build/gen/go/mfridman/debug-gateway/grpc-ecosystem/gateway/v2/service1/v1/service/service1v1gateway"
import "buf.build/gen/go/mfridman/debug-gateway/grpc-ecosystem/gateway/v2/service1/v1/service1v1gateway"

func main() {
	service1v1gateway.RegisterServiceHandler(nil, nil, nil)
}
