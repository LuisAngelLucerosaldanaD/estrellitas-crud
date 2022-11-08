
-- +migrate Up
CREATE TABLE IF NOT EXISTS auth.users(
    id uuid NOT NULL PRIMARY KEY,
    name VARCHAR (100) NOT NULL,
    lastname VARCHAR (100) NOT NULL,
    email VARCHAR (100) NOT NULL,
    birth_date VARCHAR (50) NOT NULL,
    sexo VARCHAR (10) NOT NULL,
    age INTEGER  NOT NULL,
    password VARCHAR (255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now()
);

-- +migrate Down
DROP TABLE IF EXISTS auth.users;
