# Products

This project was created following clean architecture

## VSCode 

Dev container extension

```
Reopen in container
```

## Docker
Build
```
docker build -t productsgo:latest -f .devcontainer/Dockerfile .
```
Run
```
docker run -p 9000:9000 productsgo:latest
```

## Inside container or package manager

```
go mod tidy
```
