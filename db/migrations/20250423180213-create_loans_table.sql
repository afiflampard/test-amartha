
-- +migrate Up
CREATE TABLE loans(
    "id" uuid NOT NULL,
    "borrower_id" uuid NOT NULL,
    "principal_amount" NUMERIC(20,2) NOT NULL,
    "rate" NUMERIC(5,2) NOT NULL,
    "roi" NUMERIC(5,2) NOT NULL,
    "agreement_letter_link" text,
    "status" text NOT NULL,
    "employee_id" uuid,
    "created_at" timestamptz NOT NULL DEFAULT now(),
    "updated_at" timestamptz,
    PRIMARY KEY("id"),
    CONSTRAINT "fk_loan_user" FOREIGN KEY ("borrower_id") REFERENCES "users"
    CONSTRAINT "fk_loan_employee" FOREIGN KEY ("employee_id") REFERENCES "users"
);

-- +migrate Down
DROP TABLE loans;
