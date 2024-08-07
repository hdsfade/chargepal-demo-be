package emailserver

import (
	"github.com/jordan-wright/email"
	"log"
	"net/smtp"
	"os"
	"time"
)

var EmailCh chan<- *email.Email

func InitEmail() {
	em_username := os.Getenv("EM_USERNAME")
	em_password := os.Getenv("EM_PASSWORD")
	EmailCh = NewPool(em_username, em_password, 4)
}

func NewPool(username, password string, count int, servers ...string) chan<- *email.Email {
	// default gmail
	address := "smtpout.secureserver.net:587"
	host := "smtpout.secureserver.net"
	if len(servers) == 2 {
		address = servers[0]
		host = servers[1]
	}

	emailCh := make(chan *email.Email, 10)
	var err error
	p, err := email.NewPool(
		address,
		count,
		smtp.PlainAuth("", username, password, host),
	)
	if err != nil {
		log.Fatal("init email pool error:", err)
	}
	for i := 0; i < count; i++ {
		go func() {
			for e := range emailCh {
				log.Println("send email:", e)
				err := p.Send(e, 10*time.Second)
				if err != nil {
					log.Println("send email error:", err)
				}
			}
		}()
	}
	return emailCh
}
