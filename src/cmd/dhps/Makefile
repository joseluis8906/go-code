.PHONY: build
build:
	@gazelle .
	@bazel build .

.PHONY: run
run:
	@bazel run .

.PHONY: debug
debug:
	@dlv debug --headless --listen=:2345 --api-version=2 -- $$(bazel info bazel-bin)/src/cmd/dhps/dhps_/dhps
