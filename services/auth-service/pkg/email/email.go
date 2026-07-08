package email

import "log"

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
	if cfg.Password == "" {
		log.Println("smtp pw kosong")
		return &Mailer{
			from: cfg.From,
		},nil
	}

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