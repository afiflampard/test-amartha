
-- +migrate Up
CREATE TABLE role (
    "id" UUID not null,
    "name" TEXT NOT NULL UNIQUE,
    "created_at" timestamptz DEFAULT now(),
    "updated_at" timestamptz,
    PRIMARY KEY("id")
);


-- +migrate Down
DROP TABLE role;