PROJECT := golang-memtest

clean:
	rm -f $(PROJECT)
	sudo docker stop $(PROJECT) >/dev/null 2>&1 | true
	sudo docker rm $(PROJECT) >/dev/null 2>&1 | true
	sudo docker image rm $(PROJECT) | true

go-compile:
	CGO_ENABLED=0 go build

docker-build: go-compile
	sudo docker build -f Dockerfile -t $(PROJECT) .

docker-rm:
	# ignore errors when line begins with '-'
	sudo docker stop $(PROJECT) >/dev/null 2>&1 | true
	sudo docker rm $(PROJECT) >/dev/null 2>&1 | true

docker-run: docker-build docker-rm
	sudo docker run -it $(PROJECT)

docker-run-ok: docker-rm
	sudo docker run -it -m 8m --memory-swap 8m -e nmb=4 $(PROJECT)

docker-run-bigmem: docker-rm
	sudo docker run -it -m 8m --memory-swap 8m -e nmb=12 $(PROJECT)

