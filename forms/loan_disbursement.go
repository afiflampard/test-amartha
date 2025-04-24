package forms

import "github.com/twinj/uuid"

type LoanDisbursementForm struct{}

type LoanDisbursementInput struct {
	LoanID uuid.UUID `form:"loan_id" json:"loan_id" binding:"required"`
}
