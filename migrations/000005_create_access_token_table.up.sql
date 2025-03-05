CREATE TABLE access_tokens (
    id SERIAL PRIMARY KEY,
    user_id integer references users(id),
    token varchar(255) not null,
    created_at date not null,
    expires_at date not null 
)