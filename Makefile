debugserver:
	go build

.PHONY: release

release:
	docker build -t sedlakf/debugserver .
	docker push sedlakf/debugserver
