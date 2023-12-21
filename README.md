# ftpdadmin
Admin Web Interface for Admin ProFTPD Server



## Build Required Dependencies
 - golang >= 1.21
 - upx
 - docker
 - docker-compose

## Build and Run

```
docker run --rm -v "$PWD":/build -w /build golang:1.21-bullseye make build-in-docker
cd dist 
./proftpdadmin
## http://localhost:3000
```

## Docker-Compose Run

```
git clone https://github.com/jniltinho/ftpdadmin.git
cd ftpdadmin
docker run --rm -v "$PWD":/build -w /build golang:1.21-bullseye make build-in-docker
mv dist install/proftpd-config/Docker/
cd install/proftpd-config/Docker/
docker-compose up -d
docker-compose ps

## http://localhost:8080
## LOGIN: admin PASS: admin
```

## Links
- [Fiber Examples](https://github.com/gofiber/recipes)
- [GoFiber](https://github.com/gofiber/fiber)
- [GoFiber Docs](https://docs.gofiber.io)
- [Gorm](https://gorm.io)
- [Gorm Docs](https://gorm.io/docs)
- [Air](https://github.com/cosmtrek/air)
- [Fiber Crud](https://eternaldev.com/blog/building-basic-crud-operations-in-go-with-fiber)