build:
	go build -o tweet .

build_linux:
	CGO_ENABLED=0 GOOS=linux go build -o tweet .

build_all:
	goreleaser --snapshot --rm-dist --skip-validate

build_release:
	goreleaser --rm-dist
