package api

import (
	"golang-rest-api-template/pkg/api/asset"
	"golang-rest-api-template/pkg/api/link"
	"golang-rest-api-template/pkg/api/reward"
	"golang-rest-api-template/pkg/api/user"
	"golang-rest-api-template/pkg/auth"
	"golang-rest-api-template/pkg/middleware"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.Use(gin.Logger())
	if gin.Mode() == gin.ReleaseMode {
		r.Use(middleware.Security())
		r.Use(middleware.Xss())
	}
	r.Use(middleware.Cors())
	r.Use(middleware.RateLimiter(rate.Every(1*time.Minute), 60)) // 60 requests per minute

	v1 := r.Group("/api/v1")
	{
		v1.GET("/users/info", middleware.JWTAuth(), user.GetUserInfo)
		v1.GET("/users/asset", middleware.JWTAuth(), user.GetUserAsset)
		v1.POST("/login", auth.LoginHandler)
		v1.POST("/register", auth.RegisterHandler)
		v1.POST("/users/point/buy", middleware.JWTAuth(), user.BuyPoint)
		v1.POST("/users/membership/buy", middleware.JWTAuth(), user.BuyMembership)

		v1.GET("/rewards/nfts/:id", reward.GetNFTRewardAndProof)
		v1.GET("/rewards/users/:address", reward.GetUserRewardAndProof)
		v1.GET("/assets/:id", asset.GetPrice)

		v1.POST("/link/verification/verify", middleware.JWTAuth(), link.VerifyCode)
		v1.POST("/link/verification/request", middleware.JWTAuth(), link.RequestVerificationCode)

	}

	return r
}
