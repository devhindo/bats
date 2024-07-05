package api

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"net/smtp"
	"os"
)


type loginAuth struct {
  username, password string
}

func LoginAuth(username, password string) smtp.Auth {
	return &loginAuth{username, password}
}

func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte(a.username), nil
}

func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		default:
			return nil, errors.New("Unkown fromServer")
		}
	}
	return nil, nil
}


// usage: 
// auth := LoginAuth("loginname", "password")
// err := smtp.SendMail(smtpServer + ":25", auth, fromAddress, toAddresses, []byte(message))
// or	
// client, err := smtp.Dial(smtpServer)
// client.Auth(LoginAuth("loginname", "password"))


func SendMail(to string, message []byte) error {

	fmt.Println(os.Getenv("MAIL_HOST") +":"+ os.Getenv("MAIL_PORT"))

	auth := LoginAuth(os.Getenv("MAIL_USER"), os.Getenv("MAIL_PASS"))
	err := smtp.SendMail(os.Getenv("MAIL_HOST") + ":" + os.Getenv("MAIL_PORT"), auth, os.Getenv("MAIL_USER"), []string{to}, message)
	if err != nil {
		fmt.Println("Error sending mail: ", err)
		return err
	}

	fmt.Println("Mail sent successfully")
	return nil
}

func constructConfirmationCodeMail(code string) []byte {

	subject := "Subject: Your Subject Here\n"
    body := "your confirmation code is: " + code
    message := []byte(subject + "\n" + body)


	return message
}

func generateConfirmationCode() string {
	a := ""

	// default length to 6
	for range 6 {
		num, err := rand.Int(rand.Reader, big.NewInt(9))
		if err != nil {
			panic("couldn't generate confirmation code")
		}
		a += num.Text(9)
	}

	return a
}

