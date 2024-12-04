DROP TABLE IF EXISTS products;

CREATE TABLE products (
    id serial not null unique,
    name varchar(255) not null,
    description varchar(255)
);