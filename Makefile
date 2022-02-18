go-build:
	go build -o app cmd/.

run: build
	./app

lint:
	go fmt ./...

vet:
	go vet ./...

build:
	docker build -t todo .

up: build
	docker-compose up -d #> /dev/null

down:
	docker-compose down

db:
	docker-compose start db

push:
	docker tag todo gar.io/{{YOUR PROJECT ID}}/todo
	docker push gar.io/{{YOUR PROJECT ID}}/todo


kill:
	docker kill todo

container:
	docker run  --rm -d -p 8080:8080 -e PORT='8080' \
		--name todo {{ YOUR PROJECT NAME}}