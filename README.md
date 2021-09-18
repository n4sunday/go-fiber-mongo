```sh
go mod init go-fiber-mongo
go get -u github.com/gofiber/fiber/v2
go get go.mongodb.org/mongo-driver/mongo
go get github.com/joho/godotenv
```
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
