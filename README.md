## Running Project

You can start a PostgreSQL container with the following command for local development:

    docker run -d --name cardea_db --restart=always -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=123456 -e POSTGRES_DB=cardea -p 5432:5432 postgres:15

psql connection command
    
    psql -d "host=localhost user=admin password=123456 port=5432 dbname=cardea"

Then start the app (assuming you are in the project root directory)

    go run .

Or you can directly run the project with docker-compose, but if you are using docker-compose to run, you need to change the database host to ``postgresql`` in the configs.json

    docker compose up -d
