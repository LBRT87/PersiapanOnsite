package email

import (
	"errors"
	"fmt"
	"log"

	"github.com/wneessen/go-mail"
	"golang.org/x/tools/go/analysis/passes/nilfunc"
	"honnef.co/go/tools/analysis/code"
)

type Config struct {
	Host string
	Port int
	Username string
	Password string
	From string
}

type Mailer struct {
	client *mail.client
	from string 
}



func NewMailer (cfg Config) (*Mailer,error) {
	// cek pw nya 
	if cfg.Password == "" {
		log.Println("smtp pw kosong")
		return &Mailer{
			from: cfg.From,
		},nil
	}
	// buat new client masukin (port,smtp,username,pw)
	client,err := mail.NewClient(cfg.Host,
		mail.WithPort(cfg.Host),
		mail.WithSMTP(mail.SMTPAuthPlain),
		mail.WithUsername(cfg.Username),
		mail.WithPassword(cfg.Password),
	)

	if err != nil {
		return nil,err
	}

	return &Mailer{client: client,from: cfg.From},nil

}

func (m *Mailer) SendOTP (to string, code string) error {
	if m.client == nil {
		log.Printf("email : %s, code : %s",to,code)
		return nil
	}

	msg := mail.NewMsg()

	if err := msg.From(m.From) ; err != nil {
		return err 
	}

	if err := msg.To(to); err != nil {
		return err 
	}

	msg.Subject("Kode verifikasi akun")
	msg.SetBodyString(mail.TypeTextHTML, otpHTML(code))
	return m.client.DialAndSend(msg)
}

func otpHTML (code string) string {
	return fmt.Sprint(
		`<div> 
			<h2> verifikasi Kode </h2>
			<div> 
				%s
			</div>
		</div>`,code)
}