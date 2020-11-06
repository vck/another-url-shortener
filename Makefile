run:
	go run main.go


build:
	go build  -ldflags="-s -w" main.go && mv main server
