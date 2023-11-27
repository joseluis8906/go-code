.PHONY: clean
clean:
	@bazel clean --expunge

.PHONY: gazelle
gazelle:
	@bazel run //:gazelle

.PHONY: update-repos
update-repos:
	@bazel run //:gazelle-update-repos

.PHONY: fix-bazel-imports
fix-bazel-imports:
	@ln -s $$(bazel info bazel-genfiles)/idl/dhps/dhpspb_go_proto_/github.com/joseluis8906/go-code/idl/dhpspb idl/dhpspb
