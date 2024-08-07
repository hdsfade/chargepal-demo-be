package reward

import (
	"github.com/gin-gonic/gin"
	"golang-rest-api-template/pkg/app"
	"golang-rest-api-template/pkg/e"
	"golang-rest-api-template/pkg/models"
	"net/http"
)

type ResponseNFTReward struct {
	NFTId string `json:"address"`
	Rew   string `json:"rew"`
	Proof string `json:"proof"`
}

type ResponseUserReward struct {
	Address string `json:"address"`
	Rew     string `json:"rew"`
	Proof   string `json:"proof"`
}

// GetNFTRewardAndProof godoc
//
//	@Summary		Get charge reward by nft id
//	@Description	Get nft reward by nft id
//	@Tags			reward
//	@Produce		json
//	@Param			id	path	string	true	"NFT ID"
//	@Success		200	{object}	ResponseNFTReward	"Successfully retrieved reward value and proof"
//	@Router			/rewards/nfts/{id} [get]
func GetNFTRewardAndProof(c *gin.Context) {
	var nftReward models.NFTReward
	var res ResponseNFTReward
	id := c.Param("id")
	res.NFTId = id
	if err := models.DB.Where("nft = ?", id).First(&nftReward).Error; err != nil {
		c.JSON(http.StatusOK, app.NewResponse(e.SUCCESS, res))
		return
	}

	res.Rew = nftReward.Reward
	res.Proof = nftReward.Proof
	c.JSON(http.StatusOK, app.NewResponse(e.SUCCESS, res))
}

// GetUserRewardAndProof godoc
//
//	@Summary		Get charge reward by address
//	@Description	Get user reward by address
//	@Tags			reward
//	@Produce		json
//	@Param			address	path	string	true	"user address"
//	@Success		200	{object}	ResponseUserReward	"Successfully retrieved reward value and proof"
//	@Router			/rewards/users/{id} [get]
func GetUserRewardAndProof(c *gin.Context) {
	var userReward models.UserReward
	var res ResponseUserReward
	address := c.Param("address")
	res.Address = address
	if err := models.DB.Where("address = ?", address).First(&userReward).Error; err != nil {
		c.JSON(http.StatusOK, app.NewResponse(e.SUCCESS, res))
		return
	}

	res.Rew = userReward.Reward
	res.Proof = userReward.Proof
	c.JSON(http.StatusOK, app.NewResponse(e.SUCCESS, res))
}
