// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: service2/v1/service.proto

/*
Package service2v1gateway is a reverse proxy.

It translates gRPC into RESTful JSON APIs.

Deprecated: This package has moved to "buf.build/gen/go/mfridman/debug-gateway/grpc-ecosystem/gateway/v2/service2/v1/service2v1gateway". Use that import path instead.
*/
package service2v1gateway

import gateway "buf.build/gen/go/mfridman/debug-gateway/grpc-ecosystem/gateway/v2/service2/v1/service2v1gateway"

var (

	// RegisterServiceHandlerServer registers the http handlers for service Service to "mux".
	// UnaryRPC     :call ServiceServer directly.
	// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
	// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterServiceHandlerFromEndpoint instead.
	RegisterServiceHandlerServer = gateway.RegisterServiceHandlerServer

	// RegisterServiceHandlerFromEndpoint is same as RegisterServiceHandler but
	// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
	RegisterServiceHandlerFromEndpoint = gateway.RegisterServiceHandlerFromEndpoint

	// RegisterServiceHandler registers the http handlers for service Service to "mux".
	// The handlers forward requests to the grpc endpoint over "conn".
	RegisterServiceHandler = gateway.RegisterServiceHandler

	// RegisterServiceHandlerClient registers the http handlers for service Service
	// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "service2v1grpc.ServiceClient".
	// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "service2v1grpc.ServiceClient"
	// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
	// "service2v1grpc.ServiceClient" to call the correct interceptors.
	RegisterServiceHandlerClient = gateway.RegisterServiceHandlerClient
)