CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email varchar(255) NOT NULL,
    password varchar(255) not null,
    createdAt datetime not null,
    updatedAt datetime not null
)