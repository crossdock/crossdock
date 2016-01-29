project := xlang


.PHONY: install
install:
	glide --version || go get github.com/Masterminds/glide
	GO15VENDOREXPERIMENT=1 glide install
	GO15VENDOREXPERIMENT=1 go build `GO15VENDOREXPERIMENT=1 glide novendor`


.PHONY: test
test:
	GO15VENDOREXPERIMENT=1 go test `GO15VENDOREXPERIMENT=1 glide novendor`


.PHONY: xlang
xlang:
	docker-compose run xlang


.PHONY: run
run:
	docker-compose build xlang
	docker-compose run xlang
