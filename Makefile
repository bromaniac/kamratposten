all: kamratposten

kamratposten: *.go
	go build -o kamratposten *.go

clean:
	rm -f kamratposten

rimraf:
	rm -fr items/ images/

test:
	go test

