package api

import (
	"crypto/rand"
	"math/big"

)

func SendMail(to string, content string) error {
	
	// send mail
	
	return nil
}

func constructConfirmationMail(code string) string {
	confirmationMail := ""
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