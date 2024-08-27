main_package_path = ./cmd/mauc
binary_name = mauc

.PHONY: default
default: test build

.PHONY: test
test:
	go test -v ./tests

.PHONY: build
build:
	go build -o $(binary_name) $(main_package_path)

.PHONY: clean
clean:
	rm -f $(binary_name)

.PHONY: run
run: build
	./$(binary_name)

.PHONY: build-run-clean
try: build
	./$(binary_name)
	$(MAKE) clean
