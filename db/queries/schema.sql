CREATE TABLE users (
  id SERIAL NOT NULL PRIMARY KEY,
  name VARCHAR(50) NOT NULL,
  email VARCHAR(255) NOT NULL UNIQUE,
  password VARCHAR(255) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE currencies (
  code CHAR(3) PRIMARY KEY,
  name TEXT NOT NULL,
  symbol TEXT
);

CREATE TYPE transaction_kind AS ENUM ('income', 'expense', 'fixed_expense');

CREATE TABLE categories (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  kind VARCHAR(50) NOT NULL CHECK (kind IN ('income', 'expense', 'fixed_expense'))
);

CREATE TABLE transactions (
  id SERIAL PRIMARY KEY,
  user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  kind transaction_kind NOT NULL,
  amount NUMERIC NOT NULL CHECK (amount > 0),
  currency_code CHAR(3) NOT NULL REFERENCES currencies(code),
  category_id BIGINT REFERENCES categories(id) ON DELETE
  SET NULL,
    description TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);