package forms

import (
	"fmt"

	"github.com/google/uuid"
)

type InvestForm struct{}

type InvestFormInput struct {
	LoanID uuid.UUID `form:"loan_id" json:"loan_id" binding:"required"`
	Amount float64   `form:"amount" json:"amount" binding:"required"`
}

func (i InvestForm) ValidateAmount(amount float64) error {
	if amount < 0 {
		return fmt.Errorf("amount must be bigger than 0")
	}
	return nil
}
