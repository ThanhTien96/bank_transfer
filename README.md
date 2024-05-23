## Setup development environment on Windows: WSL2, Go, VSCode, Docker, Make, Sqlc
<p dir="auto"><a href="https://github.com/stretchr/testify/actions/workflows/main.yml"><img src="https://github.com/stretchr/testify/actions/workflows/main.yml/badge.svg?branch=master" alt="Build Status" style="max-width: 100%;"></a> <a href="https://goreportcard.com/report/github.com/stretchr/testify" rel="nofollow"><img src="https://camo.githubusercontent.com/fdfe126e2cd3288499dbeb861b0a9537fb2a218878b1aa3b2a6897f30f3367c7/68747470733a2f2f676f7265706f7274636172642e636f6d2f62616467652f6769746875622e636f6d2f73747265746368722f74657374696679" alt="Go Report Card" data-canonical-src="https://goreportcard.com/badge/github.com/stretchr/testify" style="max-width: 100%;"></a> <a href="https://pkg.go.dev/github.com/stretchr/testify" rel="nofollow"><img src="https://camo.githubusercontent.com/e54fb7d8d00d4be0ccd81b83495761cee4d7de1595045f5e1402b76b893589ea/68747470733a2f2f706b672e676f2e6465762f62616467652f6769746875622e636f6d2f73747265746368722f74657374696679" alt="PkgGoDev" data-canonical-src="https://pkg.go.dev/badge/github.com/stretchr/testify" style="max-width: 100%;"></a></p>

#### 1. Install Golang on Linux

```bash
$ sudo snap install go --classic
```

#### 2. Install sqlc

```bash
$ sudo snap install sqlc
```

#### 3. Install make

```bash
$ sudo apt install make
```
<br>

## Register postgres container name postgres12

```bash
$ sudo docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine
```
<br>


## Write Migrate 

### Install [go-migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate#linux-deb-package)

```bash
$ curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add -
$ echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ $(lsb_release -sc) main" > /etc/apt/sources.list.d/migrate.list
$ apt-get update
$ apt-get install -y migrate
```
### Migrate syntax

```bash
$ migrate create -ext sql -dir db/migration -seq init_schema
```
<br>

## Make File look like
```
postgres:
	sudo docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

rmpostgres:
	sudo docker container stop postgres12
	sudo docker container rm postgres12
	sudo docker image rm postgres:12-alpine

createdb:
	sudo docker exec -it postgres12 createdb --username=root --owner=root simple_bank

dropdb:
	sudo docker exec -it postgres12 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc: 
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres rmpostgres createdb dropdb migratedown migrateup sqlc test
```

## First time migrate table into Database. 

- if you don't have table in your db yet you run this CLI
```bash
#generate table with exists query in the code
$ make migrateup

#if you want to clear the table in database immediately
$ make migratedown
```

<hr>

### Mock db testing

- First install gomock package https://github.com/golang/mock 
- Make it is executeale from anywhere.

```bash 
	vim ~/.bashrc 
	#or 
	vim ~/.bash_profile

	#add this one
	export PATH=$PATH:~/go/bin
```

```bash
$ mockgen -package mockdb -destination db/mock/Store.go simplebank/db/sqlc Store
```