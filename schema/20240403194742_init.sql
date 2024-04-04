-- +goose Up
-- +goose StatementBegin
CREATE TABLE owner (
    id uuid PRIMARY KEY,
    name varchar(32) not null,
    surname varchar(32) not null,
    patronymic varchar(32) not null
);
CREATE TABLE cars (
    registration_number VARCHAR(16) PRIMARY KEY,
    mark varchar(64) not null,
    model varchar(64) not null,
    year int not null,
    owner_id uuid references owner(id)
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE cars;
DROP TABLE owner;
-- +goose StatementEnd
