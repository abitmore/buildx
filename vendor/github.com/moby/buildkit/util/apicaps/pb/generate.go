package moby_buildkit_v1_apicaps //nolint:revive

//go:generate protoc -I=. -I=../../../vendor/ -I=../../../../../../ --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative,require_unimplemented_servers=false:. caps.proto
