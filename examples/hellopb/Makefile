build:
	protoc -I. --go-grpc_out=. hello.proto
	protoc -I. --go_out=. hello.proto
	mv moul.io/adapterkit/examples/hellopb/*.go .
	find . -type f
