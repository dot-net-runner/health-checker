package smtp

import (
	"health-checker/internal/services/notification"
	"net/smtp"
)

type (
	Service struct {
		config Configuration
	}

	Configuration struct {
		Host       string
		Port       string
		Login      string
		Password   string
		Sender     string
		Recipients []string
	}
)

func NewService(c Configuration) *Service {
	s := &Service{
		config: c,
	}

	return s
}

func (s *Service) Notify(m notification.Message) error {
	auth := smtp.PlainAuth("", s.config.Login, s.config.Password, s.config.Host)
	body := "Subject:" + m.Title + "\n" + m.Body

	err := smtp.SendMail(s.config.Host+":"+s.config.Port,
		auth,
		s.config.Sender,
		s.config.Recipients,
		[]byte(body))

	if err != nil {
		return err
	}

	return nil
}
