package listener

import (
	"github.com/ethereum/go-ethereum/common"
	"golang-rest-api-template/pkg/models"
	"golang-rest-api-template/pkg/tools/abi"
	"log"
	"os"
	"strings"
)

const (
	TopUp string = "TopUp"
)

func listenTopUp() {
	log.Println("Listening for top up requests...")
	chargeTokenAddress := os.Getenv("CHARGE_TOKEN_ADDRESS")
	chargeToken, err := abi.NewChargepalToken(common.HexToAddress(chargeTokenAddress), Conn)
	if err != nil {
		log.Fatal("Failed to instantiate charge token contract:", err)
	}

	registerChan := make(chan *abi.ChargepalTokenTopUp)
	sub, err := chargeToken.WatchTopUp(nil, registerChan, nil)
	if err != nil {
		log.Fatal("Failed to watch purchase topup event:", err)
	}
	defer sub.Unsubscribe()

	for {
		select {
		case event := <-registerChan:
			log.Println("Received charge token top up request:", event)
			userId, err := models.GetUserIdByAddress(strings.ToLower(event.Sender.String()))
			if err != nil {
				log.Println("Failed to get user id:", err)
			}
			if err = models.AddCharge(userId, event.Amount); err != nil {
				log.Println("Failed to add charge:", err)
			}
		}
	}
}
