package models

import (
	"boilerplate/forms"
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LoanMutation interface {
	CreateLoan(ctx context.Context, forms forms.LoanFormInput, userID uuid.UUID, filePath string) (*uuid.UUID, error)
	ApprovedLoan(ctx context.Context, loanID, userApprovedID uuid.UUID, filepath string) (*uuid.UUID, error)
	GetAllLoans(ctx context.Context, status []string) ([]Loans, error)
	GetLoansByID(ctx context.Context, id uuid.UUID) (Loans, error)
	CreateLoanInvestment(ctx context.Context, forms forms.InvestFormInput, userID uuid.UUID) (*uuid.UUID, error)
	DisbursementLoan(ctx context.Context, forms forms.LoanDisbursementInput, userID uuid.UUID, filePath string) (*uuid.UUID, error)

	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}

type gormMutation struct {
	tx *gorm.DB
}

func NewGormMutation(ctx context.Context, db *gorm.DB) LoanMutation {
	tx := db.WithContext(ctx).Begin()

	return &gormMutation{
		tx: tx,
	}
}

func (g *gormMutation) ApprovedLoan(ctx context.Context, loanID uuid.UUID, userApprovedID uuid.UUID, filepath string) (*uuid.UUID, error) {
	var (
		loan         Loans
		approvalLoan LoansApproval
	)
	if err := g.tx.First(&loan, loanID).Error; err != nil {
		return nil, err
	}
	if loan.Status != LoanProposed {
		return nil, fmt.Errorf("loan status is not proposed")
	}
	loan.UpdateLoan(LoanApproved)
	approvalLoan.CreateLoansApproval(loan, userApprovedID, filepath)
	if err := g.tx.Create(&approvalLoan).Error; err != nil {
		return nil, err
	}
	if err := g.tx.Save(&loan).Error; err != nil {
		return nil, err
	}

	return &approvalLoan.ID, nil
}

func (g *gormMutation) CreateLoan(ctx context.Context, forms forms.LoanFormInput, userID uuid.UUID, filePath string) (*uuid.UUID, error) {
	var (
		count int64
	)

	if err := g.tx.Model(&Loans{}).Where("borrower_id = ? And status IN (?)", userID, []string{LoanApproved, LoanInvested, LoanProposed}).Count(&count).Error; err != nil {
		return nil, err
	}

	if count > 0 {
		return nil, errors.New("error Borrower Already Exist in Status Proposed, Invested, and Approved")
	}

	loan := Loans{}
	loan.CreateLoan(forms, userID, filePath)

	if err := g.tx.Create(&loan).Error; err != nil {
		return nil, err
	}
	return &loan.ID, nil

}

func (g *gormMutation) GetAllLoans(ctx context.Context, status []string) ([]Loans, error) {
	var loans []Loans
	if err := g.tx.Where("status IN (?)", status).Find(&loans).Error; err != nil {
		return []Loans{}, err
	}
	return loans, nil
}

func (g *gormMutation) GetLoansByID(ctx context.Context, id uuid.UUID) (Loans, error) {
	var loan Loans
	if err := g.tx.First(&loan, id).Error; err != nil {
		return Loans{}, err
	}

	return loan, nil
}

func (g *gormMutation) CreateLoanInvestment(ctx context.Context, forms forms.InvestFormInput, userID uuid.UUID) (*uuid.UUID, error) {
	var (
		loan           Loans
		firstPrice     = forms.Amount
		loanInvestment LoansInvestment
	)
	if err := g.tx.Preload("LoansInvestment").First(&loan, forms.LoanID).Error; err != nil {
		return nil, err
	}

	if loan.Status == LoanApproved || loan.Status == LoanDisbursed {
		return nil, fmt.Errorf("loan is not proposed or invest")
	}
	for _, investmentLoan := range loan.LoansInvestment {
		firstPrice += investmentLoan.Amount
	}

	if loan.PrincipalAmount < firstPrice {
		return nil, fmt.Errorf("loan cannot more than Principal amount")
	}

	loanInvestment.CreateLoansInvestment(loan, userID, forms.Amount)
	loan.UpdateLoan(LoanInvested)

	if err := g.tx.Create(&loanInvestment).Error; err != nil {
		return nil, err
	}

	if err := g.tx.Save(&loan).Error; err != nil {
		return nil, err
	}

	return &loanInvestment.ID, nil
}

func (g *gormMutation) DisbursementLoan(ctx context.Context, forms forms.LoanDisbursementInput, userID uuid.UUID, filePath string) (*uuid.UUID, error) {
	var (
		loans                  Loans
		disbursementLoan       LoanDisbursement
		loanInvestorReturnList []LoanInvestorReturns
		loanRepaymentSummary   LoanRepaymentsSummary
		loanRepaymentsInput    LoanRepaymentsInput
	)

	if err := g.tx.Preload("LoansInvestment").First(&loans, forms.LoanID).Error; err != nil {
		return nil, err
	}

	if loans.Status != LoanInvested {
		return nil, fmt.Errorf("loan is not invested")
	}

	disbursementLoan.CreateNewLoanDisbursement(userID, filePath)
	for _, investor := range loans.LoansInvestment {
		loanReturnInput := LoanInvestorReturnsInput{
			LoanID:         loans.ID,
			InvestorID:     investor.InvestorID,
			InvestedAmount: investor.Amount,
			ReturnAmount:   (investor.Amount * (float64(loans.Roi) / 100)) + investor.Amount,
			Interest:       (investor.Amount * (float64(loans.Roi) / 100)),
		}
		var loanInvestorReturn LoanInvestorReturns
		loanInvestorReturn.CreateLoanInvestorReturns(loanReturnInput)
		loanInvestorReturnList = append(loanInvestorReturnList, loanInvestorReturn)
	}

	loanRepaymentsInput = LoanRepaymentsInput{
		LoanID:                 loans.ID,
		TotalPayableByBorrower: (loans.PrincipalAmount*(float64(loans.Rate)/100) + loans.PrincipalAmount),
		TotalInterest:          (loans.PrincipalAmount * (float64(loans.Rate) / 100)),
	}

	loanRepaymentSummary.CreateLoanRepaymentsSummary(loanRepaymentsInput)
	loans.UpdateLoan(LoanDisbursed)

	if err := g.tx.Create(&loanInvestorReturnList).Error; err != nil {
		return nil, err
	}

	if err := g.tx.Create(&disbursementLoan).Error; err != nil {
		return nil, err
	}

	if err := g.tx.Create(&loanRepaymentSummary).Error; err != nil {
		return nil, err
	}

	if err := g.tx.Save(&loans).Error; err != nil {
		return nil, err
	}

	return &disbursementLoan.ID, nil

}

func (g *gormMutation) Commit(ctx context.Context) error {
	return g.tx.Commit().Error
}

func (g *gormMutation) Rollback(ctx context.Context) error {
	return g.tx.Rollback().Error
}
