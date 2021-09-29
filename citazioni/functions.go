package citazioni

import (
	"fmt"
	"log"
	"net/smtp"

	"github.com/dennybiasiolli/go-dennybiasiolli-api/common"
)

func SendMailOnQuoteAdded(citazione Citazione) {
	if common.EMAIL_HOST_USER == "" ||
		common.EMAIL_HOST_PASSWORD == "" ||
		common.EMAIL_HOST == "" ||
		common.EMAIL_PORT == "" ||
		common.EMAIL_DEFAULT_FROM == "" ||
		common.EMAIL_DEFAULT_TO == "" {
		return
	}

	auth := smtp.PlainAuth("", common.EMAIL_HOST_USER, common.EMAIL_HOST_PASSWORD, common.EMAIL_HOST)
	from := common.EMAIL_DEFAULT_FROM
	to := []string{common.EMAIL_DEFAULT_TO}
	msg := []byte("From: " + from + "\r\n" +
		"To: " + common.EMAIL_DEFAULT_TO + "\r\n" +
		"Subject: New quote added - dennybiasiolli.com API Go\r\n" +
		"MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n" +
		"\r\n" +
		"<html><body>" +
		"<p>New quote added by a guest</p>" +
		"<p>ID: " + fmt.Sprintf("%v", citazione.ID) + "</p>" +
		"<p><i>" + citazione.Frase + "</i></p>" +
		"<p>Author: " + citazione.Autore + "</p>" +
		"</body></html>")
	err := smtp.SendMail(common.EMAIL_HOST+":"+common.EMAIL_PORT, auth, from, to, msg)
	if err != nil {
		log.Println(err)
	}
}
