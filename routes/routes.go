package routes

import (
	"boilerplate/controllers"
	"boilerplate/db"
	"boilerplate/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthorizeRoleMiddleware(requiredRole []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		existingRole := false
		for _, roleRequired := range requiredRole {
			if role == roleRequired {
				existingRole = true
				break
			}
		}

		if !exists || !existingRole {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Forbidden: insufficient role"})
			return
		}
		c.Next()
	}
}

func Routes(router *gin.Engine) {

	v1 := router.Group("/v1")
	{
		user := controllers.NewUserServiceMutation(db.GetDB())

		v1.POST("/user/login", user.Login)
		v1.POST("/user/register", user.Register)

		loan := controllers.NewLoanServiceMutation(db.GetDB())

		roleLoan := v1.Group("/loan")
		roleLoan.POST("/create", middleware.AuthMiddleware(), AuthorizeRoleMiddleware([]string{"borrower", "investor", "employee"}), loan.Loans)
		roleLoan.POST("/approved", middleware.AuthMiddleware(), AuthorizeRoleMiddleware([]string{"employee"}), loan.ApprovedByEmployee)
		roleLoan.POST("/invested", middleware.AuthMiddleware(), AuthorizeRoleMiddleware([]string{"investor"}), loan.LoanInvestment)
		roleLoan.POST("/disbursed", middleware.AuthMiddleware(), AuthorizeRoleMiddleware([]string{"employee"}), loan.LoanDisbursement)

		roleLoan.GET("/loan-list", middleware.AuthMiddleware(), AuthorizeRoleMiddleware([]string{"borrower", "investor", "employee"}), loan.GetAllLoans)
		roleLoan.GET("/:id", middleware.AuthMiddleware(), AuthorizeRoleMiddleware([]string{"borrower", "investor", "employee"}), loan.GetLoanByID)

	}
}
