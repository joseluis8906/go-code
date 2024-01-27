BAZEL_BIN = $$(bazel info bazel-bin)

.PHONY: build
build:
	@gazelle .
	@bazel build //src/playground/cmd/example1

.PHONY: clean
clean:
	@bazel clean

.PHONY: run
run:
	@cd /$(BAZEL_BIN)/src/playground/cmd/example1/example1_/ && ./example1

.PHONY: test 
test:
	@gazelle .
	@bazel test //src/playground/...

.PHONY: debug
debug:
	@dlv debug --headless --listen=:2345 --api-version=2 -- $$(bazel info bazel-bin)/src/cmd/example1/example1_/example1
