# Products

This project was created following clean architecture.
I choose this architecture because it has advantages like scalability, maintainability and modularity and follows SOLID principles.

I created this project in clean architecture because i would like to show you that i can create software with clean architecture.
I'm also familiar with other design patterns such as MVC, command/query, layered architecture.

The choice of a design pattern depends on the project's type and requirements.
Maybe clean architecture is too much for a small project like this but if you would like i can create the same project in MVC

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

## Inside container or in your OS

```
go mod tidy
```
