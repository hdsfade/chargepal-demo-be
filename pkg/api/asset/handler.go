package asset

import (
	"github.com/gin-gonic/gin"
	"golang-rest-api-template/pkg/app"
	"golang-rest-api-template/pkg/e"
	"golang-rest-api-template/pkg/models"
	"net/http"
)

type ResponseAssetPrice struct {
	Id    string `json:"id"`
	Price string `json:"price"`
}

//	@BasePath		/api/v1
//
// GetPrice godoc
//
//	@Summary		Get asset's price
//	@Description	Get asset's price
//	@Tags			asset
//	@Security		JWTAuth
//	@Produce		json
//	@Param			id				path	string	true	"asset id"
//	@Success		200				{object} ResponseAssetPrice	"Successfully retrieved asset's price"
//	@Router			/assets/{id} [get]
func GetPrice(c *gin.Context) {
	id := c.Param("id")
	var asset models.Asset

	if err := models.DB.Where("id = ?", id).First(&asset).Error; err != nil {
		c.JSON(http.StatusBadRequest, app.NewResponse(e.ERROR_NOT_EXIST_USER, nil))
		return
	}

	res := ResponseAssetPrice{
		Id:    asset.ID,
		Price: asset.Price,
	}

	c.JSON(http.StatusOK, app.NewResponse(e.SUCCESS, res))
}
