-- +goose Up
CREATE TABLE profiles (
    id serial primary key,
    name varchar(50),
    username varchar(50),
    email varchar(50),
    image_url text,
    phone varchar(15),
    source varchar(50),
    url text,
    location varchar(20),
    description text,
    external_id bigint,
    created_at timestamp,
    updated_at timestamp,
    birthdate timestamp,
    CONSTRAINT external_id_source UNIQUE(external_id, source)
);
-- +goose Down
DROP TABLE profiles;
