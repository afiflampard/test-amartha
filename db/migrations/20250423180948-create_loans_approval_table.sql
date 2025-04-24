
-- +migrate Up

CREATE TABLE loans_approval(
    "id" uuid NOT NULL,
    "loan_id" uuid NOT NULL,
    "field_validator_employee_id" uuid NOT NULL,
    "proof_picture_url" text,
    "approval_date" timestamptz NOT NULL,
    "created_at" timestamptz DEFAULT now(),
    "updated_at" timestamptz,
    PRIMARY KEY("id"),
    CONSTRAINT "fk_loans_approval_loans" FOREIGN KEY ("loan_id") REFERENCES "loans",
    CONSTRAINT "fk_loans_approval_id" FOREIGN KEY ("field_validator_employee_id") REFERENCES "users"
);

-- +migrate Down
DROP TABLE loans_approval;