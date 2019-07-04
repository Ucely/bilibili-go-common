package dao

import (
	"go-common/app/service/main/account-recovery/conf"
	"go-common/library/log"

	"gopkg.in/gomail.v2"
)

// SendMail send the email.
func (d *Dao) SendMail(body string, subject string, send ...string) (err error) {
	log.Info("send mail send:%v", send)
	msg := gomail.NewMessage()
	msg.SetHeader("From", conf.Conf.MailConf.Username)
	msg.SetHeader("To", send...)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", body, gomail.SetPartEncoding(gomail.Base64))
	if err = d.email.DialAndSend(msg); err != nil {
		log.Error("s.email.DialAndSend error(%v)", err)
		return
	}
	return
}
