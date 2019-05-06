cmd: protoc --go_out=plugins=grpc:. *.proto
reference: https://www.jianshu.com/p/0d5f2f561aab

mac error when use protoc: protoc-gen-go: Plugin failed with status code 1.
solution-->1. brew install protobuf. 2. copy protoc-gen-go from bin folder from go workspace to /usr/local/bin/

