COVER_PROFILE ?= coverage.out

pre-commit:
	pre-commit run --all-files

test-cover:
	go test -coverprofile=${COVER_PROFILE} ./... -v
