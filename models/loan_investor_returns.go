package models

import (
	"time"

	"github.com/google/uuid"
)

type LoanInvestorReturns struct {
	ID             uuid.UUID `gorm:"column:id" json:"id"`
	LoanID         uuid.UUID `gorm:"column:loan_id" json:"loan_id"`
	Loan           Loans     `gorm:"foreignKey:LoanID;references:ID" json:"loan"`
	InvestorID     uuid.UUID `gorm:"column:investor_id" json:"investor_id"`
	Investor       User      `gorm:"foreignKey:InvestorID;references:ID" json:"investor"`
	InvestedAmount float64   `gorm:"column:invested_amount" json:"invested_amount"`
	ReturnAmount   float64   `gorm:"column:return_amount" json:"return_amount"`
	Interest       float64   `gorm:"column:interest" json:"interest"`
	CreatedAt      time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (lir LoanInvestorReturns) TableName() string {
	return "loan_investor_returns"
}

type LoanInvestorReturnsInput struct {
	LoanID         uuid.UUID
	InvestorID     uuid.UUID
	InvestedAmount float64
	ReturnAmount   float64
	Interest       float64
}

func (lir *LoanInvestorReturns) CreateLoanInvestorReturns(input LoanInvestorReturnsInput) {
	lir.ID = uuid.New()
	lir.LoanID = input.LoanID
	lir.InvestorID = input.InvestorID
	lir.InvestedAmount = input.InvestedAmount
	lir.ReturnAmount = input.ReturnAmount
	lir.Interest = input.Interest
	lir.CreatedAt = time.Now()
	lir.UpdatedAt = time.Now()
}
