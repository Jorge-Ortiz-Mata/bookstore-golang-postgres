CREATE TABLE users (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  login VARCHAR(256) NOT NULL UNIQUE,
  password VARCHAR(256) NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL
);