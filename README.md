## Running Project

You can start a PostgreSQL container with the following command for local development:

    docker run -d --name cardea_db --restart=always -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=123456 -e POSTGRES_DB=cardea -p 5432:5432 postgres:15

Then start the app

    go run main.go

Or you can directly run the project with docker-compose:

    docker compose up -d
