all: build push clean

TAG=
ifdef tag
	TAG=$(tag)
else
	TAG=latest
endif

build: 
	CGO_ENABLED=0 GOOS=linux go build -o main main.go
	docker build -t yfsoftcom/gotest:${TAG} .

push:
	docker push yfsoftcom/gotest:${TAG}

clean:
	rm -rf main
