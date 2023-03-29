mocks:
	go install github.com/golang/mock/mockgen@v1.6.0
	go get github.com/golang/mock/gomock
	go generate ./...

tests:
	go test ./...

benchmarks:
	go test ./... -bench=. -run=XXX
