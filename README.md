# Bookstore API - Golang and Postgres

This is a basic example on how to create REST API using Golang and PostgreSQL.

## Local Development

There are some configurations you need to do before launching this web application.

10. **Install Postgres**. First, you need to have postgres installed. You can use the postgres username and password or create a new role with the required permissions. You can use the following command:

```bash
CREATE ROLE jorge_test LOGIN PASSWORD 'jorge_test' SUPERUSER;
```

20. **Add environment variables**. Open this project using VSCode or any other IDE and create a .env file in the root directory. The environment variables this project use are located in the .env.example file. Copy and paste all of them. Don't forget to change the PostgreSQL configuration according to your Postgres server. 

30. **Create and migrate DB**. Once you have access to Postgres, you can create the database and tables using the following steps

- Test DB connection: `make test-db-conn`
- Create DB for development: `make create-db-development`
- Create DB for test (for testing purposes): `make create-db-test`
- Create tables for development: `make migrate-db-development`
- Create tables for test (for testing purposes): `make migrate-db-test`

The tables definition can be found within the **dbutils/migrations** folder.

40. **Launch Application**. Once you have everything running, you can launch the web application: `make run`

## API Features

This API works under the port 8000.Within this API you can do:

- Get all books `GET http://localhost:8000/api/v1/books`
- Get a single book `GET http://localhost:8000/api/v1/books/:id`
- Create a book `POST http://localhost:8000/api/v1/books`
- Update a book `PATCH http://localhost:8000/api/v1/books/:id`
- Delete a book `DELETE http://localhost:8000/api/v1/books/:id`

## Author

- Jorge Ortiz 
- ortiz.mata.jorge@gmail.com
- San Luis Potosi S.L.P.