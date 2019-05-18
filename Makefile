protofiles:
	# Go
	protoc -I/usr/local/include -I. \
  		-I ${GOPATH}/src \
 		-I ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
 		--go_out=google/api/annotations.proto=github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api,plugins=grpc:. \
  		proto/*.proto
	# HTTP Gateway
	protoc -I/usr/local/include -I. \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--grpc-gateway_out=logtostderr=true:. \
		proto/*.proto

run-local:
	GRPC_PORT=50051 HTTP_PORT=8080 go run cmd/api/api.go
