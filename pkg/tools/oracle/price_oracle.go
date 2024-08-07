package oracle

import (
	"encoding/json"
	"errors"
	"fmt"
	"golang-rest-api-template/pkg/models"
	"log"
	"math/big"
	"net/http"
	"os"
	"time"
)

const (
	ETH        = "1027"
	retryCount = 3
)

var api = os.Getenv("ORACLE_API")
var key = os.Getenv("ORACLE_KEY")

type Response struct {
	Data map[string]struct {
		Quote struct {
			USD struct {
				Price float64 `json:"price"`
			} `json:"USD"`
		} `json:"quote"`
	} `json:"data"`
}

func StartPriceOracle() {
	go func() {
		for {
			price, err := retryGetPrice(ETH, retryCount)
			if err != nil {
				log.Println(err)
			}

			err = UpdatePrice(ETH, price)
			if err != nil {
				log.Println(err)
			}

			time.Sleep(60 * time.Second)
		}
	}()
}

func UpdatePrice(id, price string) error {
	asset := models.Asset{
		ID:    id,
		Price: price,
	}
	if err := models.DB.Save(&asset).Error; err != nil {
		return err
	}
	return nil
}

func retryGetPrice(id string, retryCount int) (string, error) {
	for i := 0; i < retryCount; i++ {
		price, err := GetPrice(id)
		if err == nil {
			return price.String(), nil
		}
		log.Println("Failed to get price:", err)
		time.Sleep(1 * time.Second)
	}
	return "0", errors.New("Failed to get price")
}

func GetPrice(id string) (*big.Int, error) {
	parameters := map[string]string{
		"id":      id,
		"convert": "USD",
	}

	headers := http.Header{
		"Accepts":           []string{"application/json"},
		"X-CMC_PRO_API_KEY": []string{key},
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest("GET", api, nil)
	if err != nil {
		return big.NewInt(0), err
	}

	req.Header = headers
	q := req.URL.Query()
	for k, v := range parameters {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return big.NewInt(0), err
	}
	defer resp.Body.Close()

	var data Response
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return big.NewInt(0), err
	}

	price := new(big.Int)
	price.SetString(fmt.Sprintf("%.0f", data.Data[id].Quote.USD.Price*1e18/1000), 10) // charge = eth/1000

	return price, nil
}
