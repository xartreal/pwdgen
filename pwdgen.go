package main

import (
	"fmt"
	"os"
        "time"

	"github.com/Luzifer/password/lib"
)

var pwd *securepassword.SecurePassword

func main() {
        pwd = securepassword.NewSecurePassword()
        for i:=1;i<10;i++ {
	password, err := pwd.GeneratePassword(12, false)
	if err != nil {
		switch {
		case err == securepassword.ErrLengthTooLow:
			fmt.Println("The password has to be more than 4 characters long to meet the security considerations")
		default:
			fmt.Println("An unknown error occured")
		}
		os.Exit(1)
	}

	fmt.Println(password)
	time.Sleep(100 * time.Millisecond)
	}
}