package link

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"golang-rest-api-template/pkg/app"
	"golang-rest-api-template/pkg/e"
	"golang-rest-api-template/pkg/models"
	"golang-rest-api-template/pkg/tools/verification"
	"net/http"
	"net/mail"
)

type VerificationCodeRequest struct {
	Email string `json: "email"`
}

type VerifyCodeRequest struct {
	Email   string `json: "email"`
	Code    string `json: "code"`
	Address string `json: "address"`
}

// RequestVerificationCode godoc
//
//	@Summary
//	@Description	request verification by email 11
//	@Tags			link
//	@Security		JwtAuth
//	@Accept			json
//	@Produce		json
//	@Param			input			body	VerificationCodeRequest	true	"verification code request"
//	@Success		200				"ok"
//	@Failure		400				{string}	string	"Bad Request"
//	@Failure		401				{string}	string	"Unauthorized"
//	@Router			/link/verification/request [post]
func RequestVerificationCode(c *gin.Context) {
	id := c.GetString("id")
	var req VerificationCodeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, app.NewResponse(e.INVALID_PARAMS, nil))
	}
	em, err := mail.ParseAddress(req.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, app.NewResponse(e.ERROR_INVALID_EMAIL, nil))
		return
	}
	verification.GenerateVerificationCodeAndSend(id, em.Address)

	c.JSON(http.StatusOK, app.NewResponse(e.SUCCESS, nil))
}

// VerifyCode godoc
//
//	@Summary		Verify and link
//	@Description	verify and link
//	@Tags			link
//	@Security		JwtAuth
//	@Accept			json
//	@Produce		json
//	@Param			input 			body	VerifyCodeRequest	true	"verify code object"
//	@Success		200				"ok"
//	@Failure		400				{string}	string	"Bad Request"
//	@Failure		401				{string}	string	"Unauthorized"
//	@Router			/link/verification/verify [post]
func VerifyCode(c *gin.Context) {
	id := c.GetString("id")
	var req VerifyCodeRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, app.NewResponse(e.INVALID_PARAMS, nil))
		return
	}
	if isSuccess := verification.VerifyCode(id, req.Email, req.Code); !isSuccess {
		c.JSON(http.StatusBadRequest, app.NewResponse(e.ERROR_INVALID_VERIFICATION, nil))
		return
	}
	if isAddress := common.IsHexAddress(req.Address); !isAddress {
		c.JSON(http.StatusBadRequest, app.NewResponse(e.ERROR_INVALID_ADDRESS, nil))
		return
	}

	var userLink models.UserLink
	if err := models.DB.Where("email = ?", req.Email).First(&userLink).Error; err != nil {
		c.JSON(http.StatusBadRequest, app.NewResponse(e.ERROR_NOT_EXIST_EMAIL, nil))
		return
	}

	var user models.User
	if err := models.DB.Where("id = ?", id).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, app.NewResponse(e.ERROR_NOT_EXIST_USER, nil))
		return
	}
	if user.Linked {
		c.JSON(http.StatusBadRequest, app.NewResponse(e.ERROR_ALREADY_LINKED, nil))
		return
	}

	user.Linked = true
	user.Email = req.Email
	user.UserID = userLink.UserID
	user.Address = req.Address

	if err := models.DB.Updates(user).Error; err != nil {
		c.JSON(http.StatusOK, app.NewResponse(e.ERROR_LINK_FAIL, nil))
		return
	}

	c.JSON(http.StatusOK, app.NewResponse(e.SUCCESS, nil))
}
