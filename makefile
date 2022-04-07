test:
	go test ./... -cover -covermode=atomic -coverprofile=coverage.txt -v

cover:
	go tool cover -html=coverage.txt
