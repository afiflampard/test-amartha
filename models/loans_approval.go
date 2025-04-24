package models

import (
	"time"

	"github.com/google/uuid"
)

type LoansApproval struct {
	ID                       uuid.UUID `gorm:"column:id" json:"id"`
	LoanID                   uuid.UUID `gorm:"column:loan_id" json:"loan_id"`
	FieldValidatorEmployeeID uuid.UUID `gorm:"column:field_validator_employee_id" json:"field_validator_employee_id"`
	FieldValidatorEmployee   User      `gorm:"foreignKey:FieldValidatorEmployeeID" json:"field_validator_employee"`
	ProofPictureUrl          string    `gorm:"column:proof_picture_url" json:"proof_picture_url"`
	ApprovalDate             time.Time `gorm:"column:approval_date" json:"approval_date"`
	CreatedAt                time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt                time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (la LoansApproval) TableName() string {
	return "loans_approval"
}

func (la *LoansApproval) CreateLoansApproval(loan Loans, userID uuid.UUID, filePath string) {
	la.ID = uuid.New()
	la.LoanID = loan.ID
	la.FieldValidatorEmployeeID = userID
	la.ProofPictureUrl = filePath
	la.ApprovalDate = time.Now()
	la.UpdatedAt = time.Now()
}
