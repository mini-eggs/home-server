default:
	sh scripts/build.sh && \
	go build && \
	./home-server

setup:
	go get github.com/pilu/fresh && \
	go get -d -v ./... && \
	go install -v ./...

dev:
	sh scripts/build.sh && \
	fresh