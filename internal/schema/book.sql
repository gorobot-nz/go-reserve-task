CREATE TABLE IF NOT EXISTS books(
    id serial PRIMARY KEY NOT NULL UNIQUE,
    title varchar(255) NOT NULL,
    authors varchar(255)[],
    book_year date
)
