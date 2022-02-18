go-build:
	go build -o app cmd/.

run: build
	./app

lint:
	go fmt ./...

vet:
	go vet ./...

build:
	docker build -t {{YOUR PROJECT NAME}} .

up: build
	docker-compose up -d #> /dev/null

down:
	docker-compose down

db:
	docker-compose start db

push:
	docker tag {{YOUR PROJECT NAME}} gar.io/{{YOUR PROJECT ID}}/{{YOUR PROJECT NAME}}
	docker push gar.io/{{YOUR PROJECT ID}}/{{YOUR PROJECT NAME}}


kill:
	docker kill {{YOUR PROJECT NAME}}

container:
	docker run  --rm -d -p 8080:8080 -e PORT='8080' \
		--name {{YOUR PROJECT NAME}} {{ YOUR PROJECT NAME}}