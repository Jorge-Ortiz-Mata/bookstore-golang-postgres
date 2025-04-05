# Bookstore API - Golang and Postgres

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