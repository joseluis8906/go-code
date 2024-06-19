.PHONY: build
build:
	@gazelle .
	@bazel build //src/mgbench/cmd/mgbench

.PHONY: build-bsd
build-bsd:
	@gazelle .
	@bazel build --platforms=@io_bazel_rules_go//go/toolchain:freebsd_amd64 //src/mgbench/cmd/mgbench
