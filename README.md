# ftpdadmin
Golang Web Interface for Admin ProFTPD



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