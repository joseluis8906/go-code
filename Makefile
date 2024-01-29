.PHONY: clean
clean:
	@bazel clean --expunge

.PHONY: gazelle
gazelle:
	@bazel run //:gazelle

.PHONY: update-repos
update-repos:
	@bazel run //:gazelle-update-repos

.PHONY: tree
tree:
	@tree -C -I scripts -I tools -I genhtml | less
