
.PHONY: generate
generate:
	cd api; buf generate

.PHONY: run
run:
	cd cmd; go run main.go

.PHONY: clean
clean:
	rm -rf pkg

.PHONY: build
build:
	cd cmd; go build main.go

.PHONY: build-flags
build-flags:
	cd cmd; go build -ldflags "-s -w" main.go
