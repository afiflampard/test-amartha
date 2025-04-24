package domain

import (
	"time"

	"github.com/google/uuid"
)

type LoanRepaymentsSummary struct {
	ID                     uuid.UUID `gorm:"column:id" json:"id"`
	LoanID                 uuid.UUID `gorm:"column:loan_id" json:"loan_id"`
	Loan                   Loans     `gorm:"foreignKey:LoanID;references:ID" json:"loan"`
	TotalPayableByBorrower float64   `gorm:"column:total_payable_by_borrower" json:"total_payable_by_borrower"`
	TotalInterest          float64   `gorm:"column:total_interest" json:"total_interest"`
	CreatedAt              time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt              time.Time `gorm:"column:updated_at" json:"updated_at"`
}

type LoanRepaymentsInput struct {
	LoanID                 uuid.UUID
	TotalPayableByBorrower float64
	TotalInterest          float64
}

func (lrs LoanRepaymentsSummary) TableName() string {
	return "loan_repayments_summary"
}

func (lrs *LoanRepaymentsSummary) CreateLoanRepaymentsSummary(input LoanRepaymentsInput) {
	lrs.ID = uuid.New()
	lrs.LoanID = input.LoanID
	lrs.TotalInterest = input.TotalInterest
	lrs.TotalPayableByBorrower = input.TotalPayableByBorrower
	lrs.CreatedAt = time.Now()
	lrs.UpdatedAt = time.Now()
}
