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
    p := persona()

	for i := 0; i < r; i++ {
		fmt.Println(text1, p)
		fmt.Println("Record Number:", i)
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

	return "\n" + fullname + "\n" + email + "\n" + sadd + "\n" + padd + "\n" + cadd + "\n" + coadd + "\n" + honum + "\n" + monum
}

func returnnum() int {

	ptr2, _ := fmt.Print(text3)
	var num int
	fmt.Scan(&num)
	if &num != nil {
		fmt.Println("Error - Invalid Number")
	} else {
		return ptr2
	}
 return num
}


