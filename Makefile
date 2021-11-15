NAME=blog
VERSION = $(shell git describe --always --tags)

all: clean
	hugo --minify
bin: clean
	hugo --minify
	statik -src=./public
	go build -o blog
build:
	hugo --minify
	statik -src=./public
	CGO_ENABLED=0 go build -o blog
up:
	supervisorctl restart blog

s: clean
	hugo server --debug -D -p 9219

clean:
	rm -rf resources public

deploy:
	docker-compose down
	docker-compose pull && docker-compose up -d
	docker system prune -f

deploybak:
	sed -i "s/version: default/version: \"$(VERSION)\"/g" deployments.yaml
	kubectl apply -f deployments.yaml
