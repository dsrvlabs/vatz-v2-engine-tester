.PHONY: proto build build_hello

build_hello:
	go build -o build/hello ./cmd/hello/main.go

proto:
	@protoc -I./proto \
		--go_out ./ --go_opt paths=import \
		--go-grpc_out ./ --go-grpc_opt paths=import \
		proto/hello/*.proto

clean:
	rm -rf ./build
