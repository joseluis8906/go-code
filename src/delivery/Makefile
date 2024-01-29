BAZEL_BIN = $$(bazel info bazel-bin)
PROJECT_NAME = delivery

.PHONY: build
build:
	@gazelle .
	@bazel build //src/$(PROJECT_NAME)/cmd/$(PROJECT_NAME)

.PHONY: test
test:
	@gazelle .
	@bazel test //src/$(PROJECT_NAME)/...

.PHONY: coverage
coverage:
	@echo $$NAME
	# @gazelle .
	# @bazel coverage //src/$(PROJECT_NAME)/$$PKG/...
	# @genhtml --no-branch-coverage -output genhtml "$(bazel info output_path)/k8-fastbuild/testlogs/src/delivery/internal/customer/customer_test/coverage.dat"

.PHONY: clean
clean:
	@bazel clean

.PHONY: build-bsd
build-bsd:
	@gazelle .
	@bazel build --platforms=@io_bazel_rules_go//go/toolchain:freebsd_amd64 //src/$(PROJECT_NAME)/cmd/$(PROJECT_NAME)

.PHONY: run
run:
	@cd /$(BAZEL_BIN)/src/$(PROJECT_NAME)/cmd/$(PROJECT_NAME)/$(PROJECT_NAME)_/ && ./$(PROJECT_NAME)

.PHONY: debug
debug:
	@dlv debug --headless --listen=:2345 --api-version=2 -- $$(bazel info bazel-bin)/src/cmd/$(PROJECT_NAME)/$(PROJECT_NAME)_/$(PROJECT_NAME)

