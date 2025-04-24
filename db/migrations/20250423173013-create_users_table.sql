
-- +migrate Up
CREATE TABLE users(
    "id" uuid NOT NULL,
    "email" text NOT NULL,
    "password" text NOT NULL,
    "name" text NOT NULL,
    "role_id" uuid NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT now(),
    "updated_at" timestamptz,
    PRIMARY KEY("id"),
    CONSTRAINT "fk_role" FOREIGN KEY ("role_id") REFERENCES "role"
); 

-- +migrate Down
DROP TABLE user;