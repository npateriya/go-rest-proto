all: build

deps:
	go get github.com/gorilla/mux
	go get github.com/go-sql-driver/mysql
	go get github.com/mattn/go-sqlite3

test:
	echo "unit test here"

build: deps
	mkdir -p output
	go build -o output/shipped .

