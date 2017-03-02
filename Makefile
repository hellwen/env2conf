all: build

build:
	go build

push:
	git add .
	git commit -am "ok"
	git push -u origin master

test:
	env2conf -type ini -file test.flat
	env2conf -type ini -file test.ini
	env2conf -type yaml -file test.yaml
