all: fmt push

fmt:
	sed -i "" 's/\ //g' *.go
	go fmt ./...

push:
	git add .
	git commit -m "add test code"
	git push
