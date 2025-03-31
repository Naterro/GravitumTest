CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    FullName varchar(255) NOT NULL
);
INSERT INTO users (FullName) VALUES ('Bill Gates');
commit;