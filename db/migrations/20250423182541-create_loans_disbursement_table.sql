
-- +migrate Up
CREATE TABLE loan_disbursement(
    "id" uuid NOT NULL,
    "loan_id" uuid NOT NULL,
    "signed_agreement_url" text,
    "disbursed_by_employee_id" uuid NOT null,
    "disbursed_date" timestamptz,
    "created_at" timestamptz DEFAULT now(),
    "updated_at" timestamptz,
    PRIMARY KEY("id"),
    CONSTRAINT "fk_loan_disbursement_loan" FOREIGN KEY ("loan_id") REFERENCES "loans",
    CONSTRAINT "fk_loan_disbursement_user" FOREIGN KEY ("disbursed_by_employee_id") REFERENCES "users"
);

-- +migrate Down
DROP TABLE loan_disbursement;
