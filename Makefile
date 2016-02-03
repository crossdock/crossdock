project := crossdock


.PHONY: install
install:
	glide --version || go get github.com/Masterminds/glide
	GO15VENDOREXPERIMENT=1 glide install
	GO15VENDOREXPERIMENT=1 go build `glide novendor`


.PHONY: test
test:
	GO15VENDOREXPERIMENT=1 go test `glide novendor`


.PHONY: crossdock
crossdock:
	docker-compose run crossdock


.PHONY: crossdock-fresh
crossdock-fresh:
	docker-compose kill
	docker-compose rm -f
	docker-compose pull
	docker-compose build
	docker-compose run crossdock


.PHONY: run
run:
	docker-compose build crossdock
	docker-compose run crossdock


.PHONY: scratch
scratch:
	GO15VENDOREXPERIMENT=1 CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main `GO15VENDOREXPERIMENT=1 glide novendor`
	docker build -f Dockerfile.scratch -t scratch-crossdock .
	docker run scratch-crossdock

