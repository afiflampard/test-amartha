package controllers

import (
	"boilerplate/forms"
	"boilerplate/models"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LoanServiceController struct {
	Mutation models.LoanMutation
}

var (
	loanServiceForm = new(forms.LoanForm)
	loanInvestment  = new(forms.InvestForm)
)

func NewLoanServiceMutation(db *gorm.DB) *LoanServiceController {
	return &LoanServiceController{
		Mutation: models.NewGormMutation(context.Background(), db),
	}
}

func (lsv LoanServiceController) Loans(c *gin.Context) {
	loanJson := c.PostForm("loan")
	var loanForm forms.LoanFormInput

	if err := json.Unmarshal([]byte(loanJson), &loanForm); err != nil {
		c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"message": err,
		})
		return
	}

	ctx := c.Request.Context()

	userID := c.GetString("user_id")
	err := loanServiceForm.ValidateRateAndRoi(loanForm.Rate, loanForm.Roi)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	file, err := c.FormFile("agreement_letter_link")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Agreement letter file is required"})
		return
	}

	dst := fmt.Sprintf("uploads/%s", file.Filename)
	if err := c.SaveUploadedFile(file, dst); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	responseModel, err := lsv.Mutation.CreateLoan(ctx, loanForm, uuid.MustParse(userID), dst)
	if err != nil {
		lsv.Mutation.Rollback(ctx)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	lsv.Mutation.Commit(ctx)

	c.JSON(http.StatusOK, gin.H{"message": "Successfully create Loan ", "Data": responseModel})

}

func (lsv LoanServiceController) ApprovedByEmployee(c *gin.Context) {
	var (
		loanApproved forms.LoanApprovedInput
		ctx          = c.Request.Context()
		userID       = c.GetString("user_id")
	)

	loanInvestmentJson := c.PostForm("loan_approved")
	if err := json.Unmarshal([]byte(loanInvestmentJson), &loanApproved); err != nil {
		c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"message": err,
		})
		return
	}

	file, err := c.FormFile("proof_picture_url")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Agreement letter file is required"})
		return
	}

	dst := fmt.Sprintf("uploads/%s", file.Filename)
	if err := c.SaveUploadedFile(file, dst); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	responseApprovedLoan, err := lsv.Mutation.ApprovedLoan(ctx, loanApproved.LoanID, uuid.MustParse(userID), dst)
	if err != nil {
		lsv.Mutation.Rollback(ctx)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	lsv.Mutation.Commit(ctx)
	c.JSON(http.StatusOK, gin.H{"message": "successfully approved loan", "Data": responseApprovedLoan})

}

func (lsv LoanServiceController) LoanInvestment(c *gin.Context) {
	var (
		loanInvestmentInput forms.InvestFormInput
		ctx                 = c.Request.Context()
		userID              = c.GetString("user_id")
	)

	loanInvestmentJson := c.PostForm("loan_investment")
	if err := json.Unmarshal([]byte(loanInvestmentJson), &loanInvestmentInput); err != nil {
		c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"message": err,
		})
		return
	}

	err := loanInvestment.ValidateAmount(loanInvestmentInput.Amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Agreement letter file is required"})
		return
	}

	responseLoanInvestment, err := lsv.Mutation.CreateLoanInvestment(ctx, loanInvestmentInput, uuid.MustParse(userID))
	if err != nil {
		lsv.Mutation.Rollback(ctx)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	lsv.Mutation.Commit(ctx)
	c.JSON(http.StatusOK, gin.H{"message": "successfully Invested loan", "Data": responseLoanInvestment})
}

func (lsv LoanServiceController) LoanDisbursement(c *gin.Context) {

	var (
		loanDisbursement forms.LoanDisbursementInput
		ctx              = c.Request.Context()
		userID           = c.GetString("user_id")
	)
	loanJsonDisbursementJson := c.PostForm("loan_disbursement")

	if err := json.Unmarshal([]byte(loanJsonDisbursementJson), &loanDisbursement); err != nil {
		c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"message": err,
		})
		return
	}

	file, err := c.FormFile("signed_agreement_url")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Agreement letter file is required"})
		return
	}

	dst := fmt.Sprintf("uploads/%s", file.Filename)
	if err := c.SaveUploadedFile(file, dst); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	responseDisbursedLoan, err := lsv.Mutation.DisbursementLoan(ctx, loanDisbursement, uuid.MustParse(userID), dst)
	if err != nil {
		lsv.Mutation.Rollback(ctx)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	lsv.Mutation.Commit(ctx)
	c.JSON(http.StatusOK, gin.H{"message": "successfully disbursed loan", "Data": responseDisbursedLoan})
}
