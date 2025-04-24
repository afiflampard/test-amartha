package forms

import (
	"fmt"

	"github.com/google/uuid"
)

type LoanForm struct{}

type LoanFormInput struct {
	PrincipalAmount float64 `form:"principal_amount" json:"principal_amount" binding:"required"`
	Rate            float32 `form:"rate" json:"rate" binding:"required"`
	Roi             float32 `form:"roi" json:"roi" binding:"required"`
}

type LoanApprovedInput struct {
	LoanID uuid.UUID `form:"loan_id" json:"loan_id"`
}

func (f LoanForm) ValidateRateAndRoi(rate, roi float32) error {
	if rate > 100 || rate < 0 {
		return fmt.Errorf("rate must be between 0 and 100, got %.2f", rate)
	}

	if roi > 100 || rate < 0 {
		return fmt.Errorf("roi must be between 0 and 100, got %.2f", rate)
	}

	return nil
}
