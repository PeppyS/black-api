protofiles:
	# Go
	protoc -I/usr/local/include -I. \
  		-I ${GOPATH}/src \
 		-I ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
 		--go_out=google/api/annotations.proto=github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api,plugins=grpc:. \
  		proto/*.proto
	# HTTP Gateway
	protoc -I/usr/local/include -I. \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--grpc-gateway_out=logtostderr=true:. \
		proto/*.proto

run-local:
	godotenv -f .env go run cmd/api/api.go

import-businesses:
	godotenv -f .env go run cmd/import/import_businesses.go

build-image:
	docker build -t xpeppy/black-api .

run-image:
	docker run -t -i --env-file .env -p 8080:8080 xpeppy/black-api

push-image:
	docker push xpeppy/black-api

deploy-image:
	docker tag xpeppy/black-api registry.heroku.com/black-api/web
	docker push registry.heroku.com/black-api/web
	heroku container:release web -a black-api
