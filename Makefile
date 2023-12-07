OUTPUT_DIR='.'
PROTO_DIR='../proto/'
COVERAGE_PATH='coverage.out'

.PHONY: proto
proto:
	@protoc --go_out=${OUTPUT_DIR} \
		--go-grpc_out=${OUTPUT_DIR} \
		--proto_path=${PROTO_DIR} \
		${PROTO_DIR}/*.proto

.PHONY: coverage-test
coverage-test:
	@go test -v ./... -coverprofile=${COVERAGE_PATH}
	@go tool cover -html=${COVERAGE_PATH}