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


.PHONY: publish
publish:
	./scripts/publish-to-docker-registry.sh
