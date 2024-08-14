
PROTO_DIR=proto
PROTO_FILE=$(PROTO_DIR)/service.proto
OUT_DIR=server
DIST_DIR=dist

generate:
	protoc -I=$(PROTO_DIR) --go_out=$(OUT_DIR) --go-grpc_out=$(OUT_DIR) $(PROTO_FILE)

build: generate
	mkdir -p $(DIST_DIR)
	go build -o $(DIST_DIR)/vendor-service ./cmd/vendor-service/main.go
	go build -o $(DIST_DIR)/product-service ./cmd/product-service/main.go

vendor:
	go run ./cmd/vendor-service/main.go

product:
	go run ./cmd/product-service/main.go

serve:
	go run ./cmd/vendor-service/main.go &
	go run ./cmd/product-service/main.go &
	wait

clean:
	rm -f $(OUT_DIR)/*.pb.go
	rm -rf $(DIST_DIR)

.PHONY: generate build run clean vendor product serve
