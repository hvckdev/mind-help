CREATE TABLE diaries
(
    id         serial primary key,
    owner_id   integer REFERENCES users (id),
    name       varchar(255) not null,
    updated_at date         not null,
    created_at date         not null
)