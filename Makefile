default: test

vendor:
	go mod vendor

install:
	go install ./...

lint:
	golint github.com/nordcloud/go-pingdom/pingdom
	golint github.com/nordcloud/go-pingdom/pingdom_ext

test:
	go test -cover github.com/nordcloud/go-pingdom/pingdom
	go test -cover github.com/nordcloud/go-pingdom/pingdom_ext

acceptance:
	PINGDOM_ACCEPTANCE=1 PINGDOM_EXT_ACCEPTANCE=1 go test github.com/nordcloud/go-pingdom/acceptance

cov:
	go test github.com/nordcloud/go-pingdom/pingdom -coverprofile=coverage.out
	go test github.com/nordcloud/go-pingdom/pingdom_ext -coverprofile=coverage.out
	go tool cover -func=coverage.out
	rm coverage.out

.PHONY: default vendor vendor_update install test acceptance cov
