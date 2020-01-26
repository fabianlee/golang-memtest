OWNER := fabianlee
PROJECT := golang-memtest
VERSION := 1.0.0
OPV := $(OWNER)/$(PROJECT):$(VERSION)

## builds docker image
docker-build:
	sudo docker build -f Dockerfile -t $(OPV) .

## cleans docker image
clean:
	sudo docker image rm $(OPV) | true

## runs container in foreground
docker-test:
	sudo docker run -it --rm $(OPV) /bin/sh

## run container in background
docker-run:
	sudo docker run -d --rm --name $(PROJECT) $(OPV)

## get to console of running container
docker-cli:
	sudo docker exec -it $(PROJECT) /bin/sh

## tails docker logs
docker-logs:
	sudo docker logs -f $(PROJECT)

## stops container running in background
docker-stop:
	sudo docker stop $(PROJECT)

## pushes to docker hub
docker-push:
	sudo docker push $(OPV)

#########################################

go-compile:
	CGO_ENABLED=0 go build

docker-test-ok:
	sudo docker run -it --rm -m 8m --memory-swap 8m -e nmb=4 $(OPV)

docker-test-bigmem:
	sudo docker run -it --rm -m 8m --memory-swap 8m -e nmb=12 $(OPV)

docker-test-bigmem-slow:
	sudo docker run -it --rm -m 8m --memory-swap 8m -e nmb=12 -e nms=2000 $(OPV)
