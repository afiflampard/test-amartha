package models

import (
	"time"

	"github.com/google/uuid"
)

type LoanDisbursement struct {
	ID                    uuid.UUID `gorm:"column:id" json:"id"`
	LoanID                uuid.UUID `gorm:"column:loan_id" json:"loan_id"`
	SignedAgreementUrl    string    `gorm:"column:signed_agreement_url" json:"signed_agreement_url"`
	DisbursedByEmployeeID uuid.UUID `gorm:"column:disbursed_by_employee_id" json:"disbursed_by_employee_id"`
	DisbursedByEmployee   User      `gorm:"foreignKey:id" json:"disbursed_by_employee"`
	DisbursedDate         time.Time `gorm:"column:disbursed_date" json:"disbursed_date"`
	CreatedAt             time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt             time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (ld LoanDisbursement) TableName() string {
	return "loans_disbursement"
}

func (lrs *LoanDisbursement) CreateNewLoanDisbursement(employeeID uuid.UUID, filePath string) {
	lrs.ID = uuid.New()
	lrs.SignedAgreementUrl = filePath
	lrs.DisbursedByEmployeeID = employeeID
	lrs.DisbursedDate = time.Now()
	lrs.CreatedAt = time.Now()
	lrs.UpdatedAt = time.Now()
}
