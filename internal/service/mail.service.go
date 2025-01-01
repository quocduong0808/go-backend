package service

import (
	"fmt"
	"go/go-backend-api/global"
	"net/smtp"

	"go.uber.org/zap"
)

type IMailService interface {
	SendTextMail(to []string, subject string, msg string) error
}

type mailService struct{}

// SendTextMail implements IMailService.
func (m *mailService) SendTextMail(toList []string, subject string, msg string) error {
	// host is address of server that the
	// sender's email address belongs,
	// in this case its gmail.
	// For e.g if your are using yahoo
	// mail change the address as smtp.mail.yahoo.com
	host := global.Config.SMTP.Host

	// Its the default port of smtp server
	port := global.Config.SMTP.Port

	// We can't send strings directly in mail,
	// strings need to be converted into slice bytes
	content := fmt.Sprintf("To: %s\r\n"+
		"Subject: %s\r\n"+
		"\r\n"+
		"%s\r\n", toList, subject, msg)
	body := []byte(content)

	// PlainAuth uses the given username and password to
	// authenticate to host and act as identity.
	// Usually identity should be the empty string,
	// to act as username.
	auth := smtp.PlainAuth("", global.Config.SMTP.Username, global.Config.SMTP.Password, host)

	// SendMail uses TLS connection to send the mail
	// The email is sent to all address in the toList,
	// the body should be of type bytes, not strings
	// This returns error if any occurred.
	err := smtp.SendMail(fmt.Sprintf("%s:%v", host, port), auth, global.Config.SMTP.Username, toList, body)

	// handling the errors
	if err != nil {
		global.Logger.Error("error while sending mail", zap.Error(err))
		return fmt.Errorf("error while sending mail: %v", err)
	}

	global.Logger.Info("successfully sent mail to all user in toList")
	return nil
}

func NewMailService() IMailService {
	return &mailService{}
}
