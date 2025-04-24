
-- +migrate Up
CREATE TABLE loan_investor_returns(
    "id" uuid NOT NULL,
    "loan_id" uuid NOT NULL,
    "investor_id" uuid NOT NULL,
    "invested_amount" DECIMAL(15,2) NOT NULL,
    "return_amount" DECIMAL(15,2) NOT NULL,
    "interest" DECIMAL(15,2) NOT NULL,
    "created_at" timestamptz DEFAULT NOW(),
    "updated_at" timestamptz,
    PRIMARY KEY("id"),
    CONSTRAINT "fk_loan_investor_return_loan" FOREIGN KEY ("loan_id") REFERENCES "loans",
    CONSTRAINT "unique_loan_investor_returns" UNIQUE ("loan_id","investor_id")
);


-- +migrate Down
DROP TABLE loan_investor_returns;
