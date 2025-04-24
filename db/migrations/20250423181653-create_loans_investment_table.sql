
-- +migrate Up
CREATE TABLE loans_investment(
    "id" uuid NOT NULL,
    "loan_id" uuid NOT NULL,
    "investor_id" uuid NOT NULL,
    "amount" NUMERIC(20,2) NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT now(),
    "updated_at" timestamptz,
    PRIMARY KEY("id"),
    CONSTRAINT "fk_loans_investment_loans" FOREIGN KEY ("loan_id") REFERENCES "loans",
    CONSTRAINT "fk_loans_investment_user" FOREIGN KEY ("investor_id") REFERENCES "users",
    CONSTRAINT "unique_loan_investor" UNIQUE ("loan_id", "investor_id")
);

-- +migrate Down
DROP TABLE loans_investment;
