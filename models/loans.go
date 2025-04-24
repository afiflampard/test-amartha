package models

import (
	"boilerplate/forms"
	"time"

	"github.com/google/uuid"
)

const (
	LoanProposed  = "proposed"
	LoanApproved  = "approved"
	LoanInvested  = "invested"
	LoanDisbursed = "disbursed"
)

type LoanModel struct{}

type Loans struct {
	ID                  uuid.UUID          `gorm:"column:id" json:"id"`
	BorrowerID          uuid.UUID          `gorm:"column:borrower_id" json:"borrower_id"`
	Borrow              User               `gorm:"foreignKey:BorrowerID;references:ID" json:"borrow"`
	PrincipalAmount     float64            `gorm:"column:principal_amount" json:"principal_amount"`
	Rate                float32            `gorm:"column:rate" json:"rate"`
	Roi                 float32            `gorm:"column:roi" json:"roi"`
	AgreementLetterLink string             `gorm:"column:agreement_letter_link" json:"agreement_letter_link"`
	Status              string             `gorm:"column:status" json:"status"`
	LoanApproval        *LoansApproval     `gorm:"foreignKey:LoanID;references:ID" json:"loan_approval"`
	LoansInvestment     []*LoansInvestment `gorm:"foreignKey:LoanID;references:ID" json:"loan_investment"`
	LoanDisbursement    *LoanDisbursement  `gorm:"foreignKey:LoanID;references:ID" json:"loan_disbursement"`
	CreatedAt           time.Time          `gorm:"column:created_at" json:"created_at"`
	UpdatedAt           time.Time          `gorm:"column:updated_at" json:"updated_at"`
}

type LoanResponse struct {
	ID uuid.UUID `json:"id"`
}

func (l Loans) TableName() string {
	return "loans"
}

func (l *Loans) CreateLoan(form forms.LoanFormInput, userID uuid.UUID, filePath string) {
	l.ID = uuid.New()
	l.BorrowerID = userID
	l.PrincipalAmount = form.PrincipalAmount
	l.Rate = form.Rate
	l.Roi = form.Roi
	l.AgreementLetterLink = filePath
	l.Status = LoanProposed
	l.UpdatedAt = time.Now()
}

func (l *Loans) UpdateLoan(status string) {
	l.Status = LoanApproved
	l.UpdatedAt = time.Now()
}
