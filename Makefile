
.PHONY: all
all: clean test build

.PHONY: clean
clean:
	rm -f bin/*

.PHONY: test
test:
	go test ./...

.PHONY: build
build:
	go build -o bin/reference_geopoi examples/reference/geopoi/reference_geopoi.go
	go build -o bin/reference_netip examples/reference/ip/reference_netip.go
	go build -o bin/validate_email examples/validate/email/validate_email.go

