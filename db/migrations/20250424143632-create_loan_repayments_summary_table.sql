
-- +migrate Up
CREATE TABLE loan_repayments_summary(
    "id" uuid NOT NULL,
    "loan_id" uuid NOT NULL REFERENCES loans(id) ON DELETE CASCADE,
    "total_payable_by_borrower" DECIMAL(15,2) NOT NULL,
    "total_interest" DECIMAL(15,2) NOT NULL,
    "created_at" timestamptz DEFAULT NOW(),
    "updated_at" timestamptz,
    PRIMARY KEY("id"),
    CONSTRAINT "fk_loan_repayments_summary_loan" FOREIGN KEY ("loan_id") REFERENCES "loans"
);


-- +migrate Down
DROP TABLE loan_repayments_summary;
