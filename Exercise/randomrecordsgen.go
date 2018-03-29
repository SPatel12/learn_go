package main

import (
	"fmt"
	"github.com/icrowley/fake"
	"strings"
	"os"
)

const (
	text1   = "****Persona Generator****"
	text2   = "-*-*-*-*-*-*-*-*-*-*-*-*-"
	text3   = "Please enter a number:"
	text4   = "Total Number of Records:"

)
//var (
//	minNum     = 1
//	maxNum     = 10
//)

func main() {

    r:= returnnum()
	    fmt.Println(text1)

    for i := 0; i < r; i++ {
    	fmt.Println(persona(), "\n")
		fmt.Println("EOF Record Number:", i + 1)
		fmt.Println(text2)
	}
	//fmt.Printf("Contains United States of America %d ", strings.Count(til, "United States of America"))
	//fmt.Printf("Contains United Kingdom %d ", strings.Count(persona(), "United Kingdom"))
	fmt.Println(text4,r)
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

	return "\n" + "NAME: " + fullname + "\n" + "Email Address: " +
		email + "\n" + "Street Address: " + sadd + "\n" + "Zip: " + padd + "\n" + "State: "+
		cadd + "\n" + "Country: " + coadd + "\n" + "H Phone: " + honum + "\n" + "Mobile: " + monum
}

func returnnum() int {
	var num int
	fmt.Print(text3)
	n, err := fmt.Scan(&num)
	if err != nil || n != 1 {
		fmt.Println(n, err)
		os.Exit(1)
	}
    return num
}


