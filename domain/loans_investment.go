package domain

import (
	"time"

	"github.com/google/uuid"
)

type LoansInvestment struct {
	ID           uuid.UUID `gorm:"column:id" json:"id"`
	LoanID       uuid.UUID `gorm:"column:loan_id" json:"loan_id"`
	Loan         Loans     `gorm:"foreignKey:LoanID;references:ID" json:"loan"`
	InvestorID   uuid.UUID `gorm:"column:investor_id" json:"investor_id"`
	InvestorUser User      `gorm:"foreignKey:InvestorID;references:ID" json:"investor_user"`
	Amount       float64   `gorm:"column:amount" json:"amount"`
	CreatedAt    time.Time `gorm:"created_at" json:"created_at"`
	UpdatedAt    time.Time `gorm:"updated_at" json:"updated_at"`
}

func (li LoansInvestment) TableName() string {
	return "loans_investment"
}

func (li *LoansInvestment) CreateLoansInvestment(loan Loans, userID uuid.UUID, loanAmount float64) {
	li.ID = uuid.New()
	li.LoanID = loan.ID
	li.InvestorID = userID
	li.Amount = loanAmount
	li.UpdatedAt = time.Now()
}
