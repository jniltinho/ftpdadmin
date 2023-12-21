# Run ProFTPd Admin Docker

Graphical User Interface for ProFTPd with MySQL and sqlite3 support


## Run

```
git clone https://github.com/jniltinho/ftpdadmin.git
cd ftpdadmin
docker run --rm -v "$PWD":/build -w /build golang:1.21-bullseye make build-in-docker
cp -aR dist install/proftpd-config/Docker/
cd install/proftpd-config/Docker/
docker-compose up -d
docker-compose exec mariadb bash -c "mysql -u root -h mariadb --password=root proftpd < /dump/tables.sql"

## http://localhost:8080
## LOGIN: admin PASS: admin
```


## Links
  
  - https://hub.docker.com/_/mariadb
  - https://github.com/wpcodevo/golang-fiber
  - https://github.com/docker/compose
  - https://codevoweb.com/create-crud-api-in-golang-using-fiber-and-gorm/
  - https://dev.to/karanpratapsingh/connecting-to-postgresql-using-gorm-24fj
