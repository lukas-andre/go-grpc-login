generate:
	cd api; buf generate

run:
	cd cmd; go run main.go

clean:
	rm -rf pkg

wire:
	cd cmd/di_container; rm -rf wire_gen.go; wire