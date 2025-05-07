check-static-check:
	which staticcheck || (go install honnef.co/go/tools/cmd/staticcheck@latest)

static-check: check-static-check
	staticcheck ./...