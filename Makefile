.PHONY: sync build test check clean

# skill.md at repo root is the source of truth.
# Go's //go:embed cannot follow symlinks, so we copy the file into
# cmd/pnt/ for the build. Run `make sync` after editing skill.md.
sync:
	cp skill.md cmd/pnt/skill.md

build: sync
	go build -o bin/pnt ./cmd/pnt

test: sync
	go test ./...

# CI guard: fail if cmd/pnt/skill.md is stale.
check:
	diff -q skill.md cmd/pnt/skill.md

clean:
	rm -rf bin dist
