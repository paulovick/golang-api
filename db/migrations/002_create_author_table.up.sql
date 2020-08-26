CREATE TABLE authors
(
    id              SERIAL NOT NULL PRIMARY KEY,
    name            TEXT
);

ALTER TABLE books ADD COLUMN author_id INTEGER REFERENCES authors(id);