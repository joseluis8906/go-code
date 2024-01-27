BAZEL_BIN = $$(bazel info bazel-bin)

.PHONY: build
build:
	@gazelle .
	@bazel build //src/pong/cmd/pong

.PHONY: clean
clean:
	@bazel clean

.PHONY: build-bsd
build-bsd:
	@gazelle .
	@bazel build --platforms=@io_bazel_rules_go//go/toolchain:freebsd_amd64 //src/pong/cmd/pong
