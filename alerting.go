package main

import (
	"log"
	"net/smtp"
)

func alerting(report string) bool {
	auth := smtp.PlainAuth("", "sfcc.report@gmail.com", "sfcc_report_pass", "smtp.gmail.com")
	to := []string{"khan.42@hotmail.fr"}
	msg := []byte(report)
	err := smtp.SendMail("smtp.gmail.com:587", auth, "sfcc_report@gmail.com", to, msg)
	if err != nil {
		log.Printf("smt error: %s", err)
		return false
	}
	return true
}
