open=xdg-open

.PHONY: test
test:
	@bazel test //src/pkg/financial:financial_test

.PHONY: cov
cov:
	@bazel coverage --combined_report=lcov //src/pkg/financial:financial_test
	@cd ../../../ && genhtml --no-branch-coverage -output genhtml "$$(bazel info output_path)/k8-fastbuild/testlogs/src/pkg/financial/financial_test/coverage.dat"
	@cd ../../../ && $(open) ./genhtml/index.html
