package main

import (
	"fmt"
	"github.com/icrowley/fake"
	"strings"
)

const (
	text1 = "****Persona Generator****"
	text2 = "-*-*-*-*-*-*-*-*-*-*-*-*-"
	text3 = "Please enter a number:"
)
func main() {

    r:= returnnum()
	    fmt.Println(text1)
    for i := 0; i < r; i++ {
		fmt.Println(persona(),"\n")
		fmt.Println("EOF Record Number:", i + 1)
		fmt.Println(text2)
	}
}

func persona() string {

	fullname := strings.ToUpper(fake.FullNameWithPrefix())
	email := fake.EmailAddress()
	sadd := fake.StreetAddress()
	padd := fake.Zip()
	cadd := fake.State()
	coadd := fake.Country()
	honum := fake.Phone()
	monum := fake.DigitsN(11)

	return "\n" + "NAME: " + fullname + "\n" + "Email Address: " + email + "\n" + "Street Address: " + sadd + "\n" + "Zip: " + padd + "\n" + "State: "+
		cadd + "\n" + "Country: " + coadd + "\n" + "H Phone: " + honum + "\n" + "Mobile: " + monum
}

func returnnum() int {
	fmt.Print(text3)
	var num int
	fmt.Scan(&num)
 return num
}


