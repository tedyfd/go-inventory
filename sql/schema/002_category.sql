-- +goose Up
CREATE TABLE category (
    id UUID PRIMARY KEY,
    name VARCHAR (50) UNIQUE NOT NULL,
    description TEXT NULL
);  

-- +goose Down
DROP TABLE category;