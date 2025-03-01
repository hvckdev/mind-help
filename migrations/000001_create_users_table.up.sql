CREATE TABLE users (
     id SERIAL PRIMARY KEY,
     name varchar(255) not null,
     email varchar(255) not null,
     password varchar(255) not null,
     registered_at date not null,
     updated_at date not null
);