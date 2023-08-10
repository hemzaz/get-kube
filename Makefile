.PHONY: dependencies build release clean

# Install dependencies
dependencies:
	go mod init github.com/hemzaz/get-kube || true
	go mod tidy
	go mod download

# Build the binary using goreleaser
build:
	goreleaser build --snapshot --rm-dist

# Release the binary
release:
	goreleaser release --snapshot --skip-publish --rm-dist

# Clean up
clean:
	rm -rf dist
