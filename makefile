build:
	go build -o bin/loadbalancer main.go

run:
	go run main.go

clean:
	rm -rf bin
