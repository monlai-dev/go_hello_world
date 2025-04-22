package controllers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"net/http"
	"strconv"
	"webapp/internal/models/request_models"
	"webapp/internal/models/response_models"
	"webapp/internal/services"
)

var Module = fx.Provide(NewAccountController)

type AccountController struct {
	accountService services.AccountServiceInterface
}

func NewAccountController(accountService services.AccountServiceInterface) *AccountController {
	return &AccountController{
		accountService: accountService,
	}
}

func (ac *AccountController) LoginHandler(c *gin.Context) {

	var req request_models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := ac.accountService.Login(req)
	if err != nil {
		c.JSON(http.StatusForbidden, responseError("Invalid email or password"))
		return
	}

	c.JSON(http.StatusOK, responseSuccess("Login successful", []interface{}{gin.H{"token": token}}))

}

func (ac *AccountController) RegisterHandler(c *gin.Context) {

	var req request_models.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdAccount, err := ac.accountService.CreateAccount(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, responseSuccess("Account created successfully", []interface{}{createdAccount}))

}

func (ac *AccountController) ListAllAccountsHandler(accountService services.AccountServiceInterface) gin.HandlerFunc {
	return func(c *gin.Context) {

		page, _ := strconv.Atoi(c.Query("page"))
		pageSize, _ := strconv.Atoi(c.Query("pageSize"))

		accounts, err := accountService.GetAllAccounts(page, pageSize)

		if err != nil {
			c.JSON(http.StatusBadRequest, responseError("Error getting accounts"))
			return
		}

		c.JSON(http.StatusOK, responseSuccess("Accounts retrieved successfully", []interface{}{accounts}))
	}
}

func (ac *AccountController) GetRandomAccountHandler(accountService services.AccountServiceInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		account, err := accountService.GetRandomAccount()
		if err != nil {
			c.JSON(http.StatusBadRequest, responseError("Error getting account"))
			return
		}
		c.JSON(http.StatusOK, responseSuccess("Account retrieved successfully", []interface{}{account}))
	}
}

func (ac *AccountController) GetAccountByIDHandler(accountService services.AccountServiceInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		account, err := accountService.GetAccountById(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, responseError("Error getting account"))
			return
		}
		c.JSON(http.StatusOK, responseSuccess("Account retrieved successfully", []interface{}{account}))
	}
}

func (ac *AccountController) GetHomelessAccountsHandler(accountService services.AccountServiceInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		accounts, err := accountService.GetAllHomelessAccounts()
		if err != nil {
			c.JSON(http.StatusBadRequest, responseError("Error retrieving accounts"))
			return
		}
		c.JSON(http.StatusOK, responseSuccess("Accounts retrieved successfully", []interface{}{accounts}))
	}
}

func (ac *AccountController) UpdateAddressHandler(accountService services.AccountServiceInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		email, _ := c.Get("email")
		var address request_models.AddressRequest
		if err := c.ShouldBindJSON(&address); err != nil {
			c.JSON(http.StatusBadRequest, responseError("Invalid address"))
			return
		}

		err := accountService.UpdateAddress(email.(string), address)
		if err != nil {
			c.JSON(http.StatusBadRequest, responseError("Error updating address"))
			return
		}
		c.JSON(http.StatusOK, responseSuccess("Address updated successfully", nil))
	}
}

func (ac *AccountController) LogoutHandler(accountService services.AccountServiceInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		token := authHeader[7:] // Extract token from "Bearer <token>"

		if err := accountService.Logout(token); err != nil {
			c.JSON(http.StatusBadRequest, responseError("Error logging out"))
			return
		}
		c.JSON(http.StatusOK, responseSuccess("Logged out successfully", nil))
	}
}

// Response Helpers
func responseError(message string) response_models.Response {
	return response_models.Response{
		ResponseCode: http.StatusBadRequest,
		Message:      message,
	}
}

func responseSuccess(message string, data []interface{}) response_models.Response {
	return response_models.Response{
		ResponseCode: http.StatusOK,
		Message:      message,
		Data:         data,
	}
}
