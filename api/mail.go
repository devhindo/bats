package api

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"net/smtp"
	"os"
)



func SendMail(to string, content string) error {

	fmt.Println(os.Getenv("MAIL_USERNAME"), os.Getenv("MAIL_PASSWORD"), os.Getenv("MAIL_HOST"), os.Getenv("MAIL_SMTP_PORT"))
	
	fmt.Println(content)
	
	err := smtp.SendMail(os.Getenv("MAIL_HOST") + ":25", mailAuth, os.Getenv("MAIL_USERNAME"), []string{to}, []byte(content))
	fmt.Println("stuck yet?")
	if err != nil {
		fmt.Println("couldn't send email. err: ", err)
	}
	
	return nil
}

func constructConfirmationCodeMail(code string) string {
	confirmationMail := ""
	confirmationMail += "your confirmation code is: " + code
	return confirmationMail
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

func initMailService() smtp.Auth {
	return LoginAuth(os.Getenv("MAIL_USERNAME"), os.Getenv("MAIL_PASSWORD"))
}

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