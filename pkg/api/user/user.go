package user

import (
	"github.com/gin-gonic/gin"
	"golang-rest-api-template/pkg/app"
	"golang-rest-api-template/pkg/e"
	"golang-rest-api-template/pkg/models"
	"golang-rest-api-template/pkg/tools/abi"
	"golang-rest-api-template/pkg/tools/oracle"
	"gorm.io/gorm"
	"math/big"
	"net/http"
)

var membershipPrice, _ = new(big.Int).SetString("10000000000000000000", 10)

type ResponseUserInfo struct {
	Id         uint   `json:"id"`
	Username   string `json:"username"`
	Linked     bool   `json:"linked"`
	Email      string `json:"email"`
	UserId     string `json:"userId"`
	Address    string `json:"address"`
	Membership bool   `json:"membership"`
}

type ResponseUserAsset struct {
	Id     uint   `json:"id"`
	Charge string `json:"charge"`
	Point  string `json:"point"`
}

type RequestBuyPoint struct {
	Amount string `json:"amount"`
}

type ResponseBuyPoint struct {
	ChargeAmount string `json:"charge_amount"`
	PointAmount  string `json:"point_amount"`
	Price        string `json:"price"`
}

//	@BasePath		/api/v1
//
// GetUserInfo godoc
//
//	@Summary		Get user info
//	@Description	Get user's id, linked, email, uuid
//	@Tags			user
//	@Security		JwtAuth
//	@Produce		json
//	@Success		200		{object}	ResponseUserInfo	"Successfully retrieved list of books"
//	@Router			/users/info [get]
func GetUserInfo(c *gin.Context) {
	idStr := c.GetString("id")

	var user models.User

	if err := models.DB.Where("id = ?", idStr).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, app.NewResponse(e.ERROR_NOT_EXIST_USER, nil))
		return
	}

	userInfoRes := ResponseUserInfo{
		Id:         user.ID,
		Username:   user.Username,
		Linked:     user.Linked,
		Email:      user.Email,
		UserId:     user.UserID,
		Address:    user.Address,
		Membership: user.Membership,
	}

	c.JSON(http.StatusOK, app.NewResponse(e.SUCCESS, userInfoRes))

	return
}

//	@BasePath		/api/v1
//
// GetUserAsset godoc
//
//	@Summary		Get user's asset
//	@Description	Get user's asset
//	@Tags			user
//	@Security		JwtAuth
//	@Produce		json
//	@Success		200			{object}	ResponseUserAsset	"Successfully retrieved user's asset"
//	@Router			/users/asset [get]
func GetUserAsset(c *gin.Context) {
	var userAsset models.UserAsset
	id := c.GetString("id")

	if err := models.DB.Where("id = ?", id).First(&userAsset).Error; err != nil {
		userAsset.ID = 49
		userAsset.Charge = "0"
		userAsset.Point = "0"
		models.DB.Save(&userAsset)
		//c.JSON(http.StatusBadRequest, app.NewResponse(e.ERROR_NOT_EXIST_USER, nil))
		//return
	}
	resUserAsset := ResponseUserAsset{
		Id:     userAsset.ID,
		Charge: userAsset.Charge,
		Point:  userAsset.Point,
	}
	c.JSON(http.StatusOK, app.NewResponse(e.SUCCESS, resUserAsset))
}

//	@BasePath		/api/v1
//
// BuyPoint godoc
//
//	@Summary		user buy point
//	@Description	user buy point
//	@Tags			user
//	@Security		JwtAuth
//	@Produce		json
//	@Param			input 			body	RequestBuyPoint	true	"buy point object"
//	@Success		200				{object}	ResponseBuyPoint	"Successfully retrieved user's point"
//	@Router			/users/point/buy [post]
func BuyPoint(c *gin.Context) {
	id := c.GetString("id")
	var asset models.Asset
	var userAsset models.UserAsset
	var req RequestBuyPoint
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, app.NewResponse(e.INVALID_PARAMS, err))
		return
	}

	if err := models.DB.Where("id = ?", oracle.ETH).First(&asset).Error; err != nil {
		c.JSON(http.StatusOK, app.NewResponse(e.ERROR_BUY_POINT_FAIL, nil))
		return
	}

	if err := models.DB.Where("id = ?", id).First(&userAsset).Error; err != nil {
		c.JSON(http.StatusBadRequest, app.NewResponse(e.ERROR_NOT_EXIST_USER, nil))
		return
	}

	userChargeAmount := new(big.Int)
	userChargeAmount.SetString(userAsset.Charge, 10)
	price, err := oracle.GetPrice(oracle.ETH) // use ETH to mock

	tmpAmount := new(big.Int)
	pointAmount := new(big.Int)
	pointAmount.SetString(userAsset.Point, 10)
	tmpAmount.SetString(req.Amount, 10)

	if userChargeAmount.Cmp(tmpAmount) < 0 {
		c.JSON(http.StatusBadRequest, app.NewResponse(e.ERROR_INSUFFICIENT_CHARGE_BALANCE, nil))
		return
	}
	userChargeAmount = userChargeAmount.Sub(userChargeAmount, tmpAmount)
	userAsset.Charge = userChargeAmount.String()

	tmpAmount.Mul(tmpAmount, price)
	tmpAmount.Div(tmpAmount, abi.Decimal)
	pointAmount.Add(pointAmount, tmpAmount)

	userAsset.Point = pointAmount.String()

	pointEvent := models.PointEvent{
		ID:    userAsset.ID,
		Price: price.String(),
		Point: tmpAmount.String(),
	}

	err = models.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Updates(&userAsset).Error; err != nil {
			return err
		}
		if err := tx.Create(&pointEvent).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		c.JSON(http.StatusOK, app.NewResponse(e.ERROR_BUY_POINT_FAIL, nil))
		return
	}
	resBuyPoint := ResponseBuyPoint{
		ChargeAmount: req.Amount,
		PointAmount:  tmpAmount.String(),
		Price:        price.String(),
	}
	c.JSON(http.StatusOK, app.NewResponse(e.SUCCESS, resBuyPoint))
}

//	@BasePath		/api/v1
//
// BuyMembership godoc
//
//	@Summary		user buy membership
//	@Description	user buy membership
//	@Tags			user
//	@Security		JwtAuth
//	@Produce		json
//	@Success		200	{string}	string	"ok"
//	@Router			/users/membership/buy [post]
func BuyMembership(c *gin.Context) {
	id := c.GetString("id")
	var user models.User
	var userAsset models.UserAsset

	if err := models.DB.Where("id = ?", id).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, app.NewResponse(e.ERROR_NOT_EXIST_USER, nil))
		return
	}
	if user.Membership {
		c.JSON(http.StatusBadRequest, app.NewResponse(e.ERROR_ALREADY_MEMBERSHIP, nil))
		return
	}

	if err := models.DB.Where("id = ?", id).First(&userAsset).Error; err != nil {
		c.JSON(http.StatusBadRequest, app.NewResponse(e.ERROR_NOT_EXIST_USER, nil))
		return
	}

	userChargeAmount := new(big.Int)
	userChargeAmount.SetString(userAsset.Charge, 10)

	if userChargeAmount.Cmp(membershipPrice) < 0 {
		c.JSON(http.StatusBadRequest, app.NewResponse(e.ERROR_INSUFFICIENT_CHARGE_BALANCE, nil))
		return
	}
	userChargeAmount = userChargeAmount.Sub(userChargeAmount, membershipPrice)
	userAsset.Charge = userChargeAmount.String()
	user.Membership = true

	buyMembershipEvent := models.BuyMembershipEvent{
		ID: userAsset.ID,
	}

	err := models.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Updates(&userAsset).Error; err != nil {
			return err
		}
		if err := tx.Updates(&user).Error; err != nil {
			return err
		}
		if err := tx.Create(&buyMembershipEvent).Error; err != nil {
		}
		return nil
	})

	if err != nil {
		c.JSON(http.StatusOK, app.NewResponse(e.ERROR_BUY_MEMBERSHIP_FAIL, nil))
		return
	}
	c.JSON(http.StatusOK, app.NewResponse(e.SUCCESS, nil))
}
