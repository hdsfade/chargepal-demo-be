package listener

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"os"
)

var Conn *ethclient.Client

func StartListenEvent() {
	rpcUrl := os.Getenv("RPC_URL")
	var err error
	if Conn, err = ethclient.Dial(rpcUrl); err != nil {
		log.Fatalln("Failed to connect eth rpc: ", err)
	}
	log.Println("Listening for top up requests...")
	go listenTopUp()
}
