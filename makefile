SRC=./main.go
DIST=./bin/main


run: build
	@echo "Running the application..."
	$(DIST)

build:
	@echo "Building the binary file"
	GOOS=darwin GOARCH=arm64 go build -o $(DIST) $(SRC)

clean:
	@echo "Cleaning the binary file ..."
	rm -rfd  $(DIST)
