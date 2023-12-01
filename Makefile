.PHONY: build clean postgres

build: bin/data-layer

clean:
	rm -rf ./bin

bin/data-layer: *.go go.mod go.sum
	go build -o bin/data-layer

postgres: 
	docker exec -i data-layer-postgres-1 psql -U postgres -c 'DROP DATABASE NextJudge'
	docker exec -i data-layer-postgres-1 psql -U postgres -c 'CREATE DATABASE NextJudge'
	docker exec -i data-layer-postgres-1 psql -U postgres nextjudge < nextjudge.sql