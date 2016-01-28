project := xlang


.PHONY: install
install:
	go build ./...


.PHONY: test
test:
	go test ./...


.PHONY: xlang
xlang:
	docker-compose run xlang


.PHONY: run
run:
	docker-compose build xlang
	docker-compose run xlang
