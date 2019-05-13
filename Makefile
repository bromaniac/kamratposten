all: kamratposten

kamratposten: *.go
	go build -o kamratposten *.go

clean:
	rm -f kamratposten

test:
	go test

