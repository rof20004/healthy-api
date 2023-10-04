CREATE TABLE IF NOT EXISTS users(
    id         text        not null,
    name       text        not null,
    age        integer     not null,
    email      text        not null,
    created_at timestamptz not null,
    primary key(id)
);