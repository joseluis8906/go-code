.PHONY: build
build:
	@gazelle .
	@bazel build .

.PHONY: build-freebsd
build-freebsd:
	@gazelle .
	@bazel build --platforms=@io_bazel_rules_go//go/toolchain:freebsd_amd64 .

.PHONY: run
run:
	@gazelle .
	@bazel run .
