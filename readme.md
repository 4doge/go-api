# REST API on Golang

### Migrations
Install `migrate` [CLI tool](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate).
To create a new migration run:
```sh
migrate create -ext sql -dir db/migrations -seq <name>
```
To run a migration you need to specify database URL:
```sh
export POSTGRESQL_URL='postgres://goapidbadmin:s3cr3t@localhost:5432/goapidb?sslmode=disable'
```
after that you can run:
```sh
migrate -database ${POSTGRESQL_URL} -path db/migrations up
```
Or if you want to revert all migrations you can run:
```sh
migrate -database ${POSTGRESQL_URL} -path db/migrations down
```