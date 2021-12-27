CREATE TABLE IF NOT EXISTS books(
    id serial PRIMARY KEY NOT NULL UNIQUE,
    title varchar(255) NOT NULL,
    authors varchar(255)[],
    book_year date,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);