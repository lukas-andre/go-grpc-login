generate:
	cd api; buf generate

run:
	cd cmd; go run main.go

clean:
	rm -rf pkg

build:
	cd cmd; go build main.go

build-flags:
	cd cmd; go build -ldflags "-s -w" main.go

wire:
	cd internal/di_container; rm -rf wire_gen.go; wire