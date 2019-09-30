install:
		go build -o bin/main parking_lot.go
run:
		go run main.go
test:
		./bin/run_functional_tests
