# ⚡Fiber

🌐 https://github.com/gofiber/fiber#readme

## 🚀 Basic Server

```go
package main

import "github.com/gofiber/fiber/v2"

func main() {
    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello Go Fiber 🚀")
    })

    app.Listen(":3000")
}

```

## 🚀 Setup Project

```sh
go mod init go-fiber-mongo
go get -u github.com/gofiber/fiber/v2
go get go.mongodb.org/mongo-driver/mongo
go get github.com/joho/godotenv
```

## 🚀 MongoDB

📄 docker-compose.yml

```yml
version: "3.7"

services:
  mongodb:
    image: mongo:latest
    container_name: mongodb
    restart: always
    environment:
      MONGO_INITDB_ROOT_USERNAME: root
      MONGO_INITDB_ROOT_PASSWORD: 123456
      MONGO_INITDB_DATABASE: sundaydb
    ports:
      - 27017:27017
```

## 🚀 Live Reload for Go

### Installation

🌐 https://github.com/cosmtrek/air

```sh
go get -u github.com/cosmtrek/air
```

```sh
# binary will be $(go env GOPATH)/bin/air
curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

# or install it into ./bin/
curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s

air -v
```

You can initialize the .air.toml configuration file to the current directory with the default settings running the following command.

```sh
air init
```

After this you can just run the air command without additional arguments and it will use the .air.toml file for configuration.

```sh
air
```

## 🚀 Use MongoDB

#### 🔥 Query Select Field
use `Decode` 
```go
var results []models.EmployeeX

cursor, err := collection.Find(ctx, filter, findOptions)

for cursor.Next(ctx) {
    var result models.EmployeeX
    cursor.Decode(&result)
    results = append(results, result)
}
```
