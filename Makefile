all: fmt push

fmt:
	go fmt ./...

push:
	git add .
	git commit -m "add test code"
	git push
