run:
	env HTTP_PORT=3000 VERSION=0.0.1 go run cmd/http/*.go

build:
	go build -o bin/http cmd/http/*.go

package:
	docker build --no-cache=true -t cncf_http:alpine -f deployment/docker/http/Dockerfile .

publish:
	docker push cncf_http:alpine