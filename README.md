# Bookstore API - Golang and Postgres

This is a basic example on how to create REST API using Golang and PostgreSQL.  

## Setup

This API works under the port 8000. Once you have clone this repo, make sure to have your postgres database running. Within this API you can do:

- Get all books `GET http://localhost:8000/api/v1/books`
- Get a single book `GET http://localhost:8000/api/v1/books/:id`
- Create a book `POST http://localhost:8000/api/v1/books`
- Update a book `PATCH http://localhost:8000/api/v1/books/:id`
- Delete a book `DELETE http://localhost:8000/api/v1/books/:id`

Run this web application by running the following command `make run` in the terminal.

## PostgreSQL setup

In order to create a new API using PostgreSQL, this is the code used to add a new role and create the database this web is connected

```bash
CREATE ROLE jorge_test LOGIN PASSWORD 'jorge_test' SUPERUSER;

CREATE DATABASE jorge_test_db
  WITH
  OWNER = jorge_test
  ENCODING = 'UTF8'
  LC_COLLATE = 'C'
  LC_CTYPE = 'C'
  TABLESPACE = pg_default
  CONNECTION LIMIT = -1
  IS_TEMPLATE = False;
```

## Author

- Jorge Ortiz 
- ortiz.mata.jorge@gmail.com
- San Luis Potosi S.L.P.