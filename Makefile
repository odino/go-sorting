run:
	go test -bench=. ./... -benchtime=1000x
test:
	go test ./...