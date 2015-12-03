package main

import (
	"fmt"
	"os"
        "time"
	"strconv"

	"github.com/Luzifer/password/lib"
)

var pwd *securepassword.SecurePassword
var plen int
var err error

func main() {
        if len(os.Args) < 2 {
                plen=12
        } else {
			plen,err=strconv.Atoi(os.Args[1])
			if err!=nil {
				plen=12
			}
        }
        pwd = securepassword.NewSecurePassword()
        for i:=1;i<10;i++ {
	password, err := pwd.GeneratePassword(plen, false)
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