BAZEL_BIN = $$(bazel info bazel-bin)

.PHONY: build
build:
	@gazelle .
	@bazel build //src/jte/cmd/jte

.PHONY: clean
clean:
	@bazel clean

.PHONY: run
run:
	@cd /$(BAZEL_BIN)/src/jte/cmd/jte/jte_/ && ./jte

.PHONY: test 
test:
	@gazelle .
	@bazel test //src/jte/...

.PHONY: debug
debug:
	@dlv debug --headless --listen=:2345 --api-version=2 -- $$(bazel info bazel-bin)/src/cmd/jte/jte_/jte
