all: build

build: clean proto master modules

master: build/bin/master

modules: \
	build/modules/echo

build/bin/master:
	go build -o build/bin/master ./src/master/

build/modules/%: src/modules/%
	mkdir -p $@
	mkdir -p $@/bin
	mkdir -p $@/conf
	mkdir -p $@/assets
	mkdir -p $</conf
	mkdir -p $</assets
	go build -o $@/bin/main ./$<
	cp -r $</conf $@/conf
	cp -r $</assets $@/assets

proto: proto_common proto_main proto_api

proto_common:
	protoc -I src/proto --go_out=src/proto_gen/ src/proto/common.proto

proto_main:
	protoc -I src/proto --go_out=src/proto_gen/ src/proto/main.proto
	protoc -I src/proto --go-grpc_out=src/proto_gen/ src/proto/main.proto

proto_api:
	protoc -I src/proto --go_out=src/proto_gen/ src/proto/api.proto
	protoc -I src/proto --go-grpc_out=src/proto_gen/ src/proto/api.proto

run: build
	MODULE_DIR=./build/modules/ ./build/bin/master

clean:
	rm -rf build

.PNONY: all build proto proto_main proto_api master modules run clean
