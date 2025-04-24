package controllers

import (
	"boilerplate/domain"
	"boilerplate/forms"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LoanServiceController struct {
	DB *gorm.DB
}

var (
	loanServiceForm = new(forms.LoanForm)
	loanInvestment  = new(forms.InvestForm)
)

func NewLoanServiceMutation(db *gorm.DB) *LoanServiceController {
	return &LoanServiceController{
		DB: db,
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
	mutation := domain.NewGormMutation(ctx, lsv.DB)
	responseModel, err := mutation.CreateLoan(ctx, loanForm, uuid.MustParse(userID), dst)
	if err != nil {
		mutation.Rollback(ctx)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	mutation.Commit(ctx)

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

	mutation := domain.NewGormMutation(ctx, lsv.DB)

	responseApprovedLoan, err := mutation.ApprovedLoan(ctx, loanApproved.LoanID, uuid.MustParse(userID), dst)
	if err != nil {
		mutation.Rollback(ctx)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	mutation.Commit(ctx)
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
	log.Println("loan input ", loanInvestmentInput)
	err := loanInvestment.ValidateAmount(loanInvestmentInput.Amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Agreement letter file is required"})
		return
	}
	mutation := domain.NewGormMutation(ctx, lsv.DB)
	responseLoanInvestment, err := mutation.CreateLoanInvestment(ctx, loanInvestmentInput, uuid.MustParse(userID))
	if err != nil {
		mutation.Rollback(ctx)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	mutation.Commit(ctx)
	c.JSON(http.StatusOK, gin.H{"message": "successfully Invested loan", "Data": responseLoanInvestment})
}

// @Summary      Disburse a loan
// @Description  Disburse a loan by employee
// @Tags         Loans
// @Accept       multipart/form-data
// @Produce      json
// @Param        loan_disbursement formData string true "LoanDisbursementInput JSON"
// @Param        signed_agreement_url formData file true "Signed Agreement File"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]interface{}
// @Router       /loan/disbursed [post]
// @Security     BearerAuth
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

	log.Println(loanDisbursement)

	dst := fmt.Sprintf("uploads/%s", file.Filename)
	if err := c.SaveUploadedFile(file, dst); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	mutation := domain.NewGormMutation(ctx, lsv.DB)
	responseDisbursedLoan, err := mutation.DisbursementLoan(ctx, loanDisbursement, uuid.MustParse(userID), dst)
	if err != nil {
		mutation.Rollback(ctx)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	mutation.Commit(ctx)
	c.JSON(http.StatusOK, gin.H{"message": "successfully disbursed loan", "Data": responseDisbursedLoan})
}

func (lsv LoanServiceController) GetLoanByID(c *gin.Context) {
	var (
		idLoan = c.Param("id")
		ctx    = c.Request.Context()
	)

	mutation := domain.NewGormMutation(ctx, lsv.DB)

	responseLoanByID, err := mutation.GetLoansByID(ctx, uuid.MustParse(idLoan))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		mutation.Rollback(ctx)
		return
	}
	mutation.Commit(ctx)
	c.JSON(http.StatusOK, gin.H{"message": "successfully get loan", "Data": responseLoanByID})
}

func (lsv LoanServiceController) GetAllLoans(c *gin.Context) {
	var (
		ctx        = c.Request.Context()
		loanStatus forms.LoanStatusInput
	)

	if err := c.ShouldBindJSON(&loanStatus); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err, "Data": nil})
		return
	}

	mutation := domain.NewGormMutation(ctx, lsv.DB)
	responseGetAllLoans, err := mutation.GetAllLoans(ctx, loanStatus.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		mutation.Rollback(ctx)
		return
	}
	mutation.Commit(ctx)
	c.JSON(http.StatusOK, gin.H{"message": "successfully get loan", "Data": responseGetAllLoans})
}
