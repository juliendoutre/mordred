.PHONY: clean build install

HASH := $(shell git rev-parse HEAD)
LDFLAGS := "-X github.com/juliendoutre/mordred/cmd/mordred/cmd.hash=${HASH}"

clean:
	rm -rf ./bin/
	rm -f c.out coverage.html

build:
	go build -o ./bin/mordred -ldflags ${LDFLAGS} ./cmd/mordred/

install:
	go install -ldflags ${LDFLAGS} ./cmd/mordred/

test:
	go test ./pkg/... -coverprofile=c.out
	go tool cover -html=c.out -o coverage.html
