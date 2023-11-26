open=xdg-open

.PHONY: test
test:
	@echo "Running tests..."
	@bazel test //src/pkg/delivery:delivery_test --test_output=errors

.PHONY: cov
cov:
	@bazel coverage --combined_report=lcov //src/pkg/delivery:delivery_test
	@cd ../../../ && genhtml --no-branch-coverage -output genhtml "$$(bazel info output_path)/k8-fastbuild/testlogs/src/pkg/delivery/delivery_test/coverage.dat"
	@cd ../../../ && $(open) ./genhtml/index.html
