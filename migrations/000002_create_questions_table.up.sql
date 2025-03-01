CREATE TABLE questions
(
    id         serial primary key,
    owner_id   integer REFERENCES users (id),
    name       varchar(255) not null,
    created_at date         not null,
    updated_at date         not null
)