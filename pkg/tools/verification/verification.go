package verification

import (
	"context"
	"github.com/jordan-wright/email"
	"golang-rest-api-template/pkg/cache"
	emailserver "golang-rest-api-template/pkg/tools/email"
	"log"
	"math/rand"
	"strconv"
	"time"
)

func GenerateVerificationCodeAndSend(id, em string) {
	code := generateVerificationCode()
	cacheKey := "id_" + id + "_em_" + em
	cache.Rdb.Set(context.TODO(), cacheKey, code, 30*time.Minute)
	sendVerificationCodeEmail(em, code)
}

func VerifyCode(id, em, code string) bool {
	cacheKey := "id_" + id + "_em_" + em
	cacheCode, err := cache.Rdb.Get(context.TODO(), cacheKey).Result()
	if err != nil || cacheCode != code {
		return false
	}
	return true
}

func generateVerificationCode() string {
	rand.NewSource(time.Now().UnixNano())
	code := ""
	for i := 0; i < 6; i++ {
		code += strconv.Itoa(rand.Intn(10))
	}
	return code
}

func sendVerificationCodeEmail(em, code string) {
	e := email.NewEmail()
	e.From = "Verification <contact@chargepal.xyz>"
	e.To = []string{em}
	e.Subject = "ChargePal Verification Code"
	log.Println("code: ", code)
	e.Text = []byte(code)
	emailserver.EmailCh <- e
}
