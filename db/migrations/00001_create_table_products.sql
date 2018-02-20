-- +goose Up
CREATE TABLE profiles (
    id serial primary key,
    username varchar(50),
    email varchar(50),
    image_url text,
    phone varchar(15),
    source varchar(50),
    url text
);
-- +goose Down
DROP TABLE profiles;
