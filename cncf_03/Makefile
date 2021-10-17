run:
	go run cmd/http/*.go

build:
	go build -o bin/http cmd/http/*.go

package:
	docker build --no-cache=true -t cncf_http:alpine -f deployment/docker/http/Dockerfile .

