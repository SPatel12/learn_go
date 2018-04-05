package main

import (
	"fmt"
	"github.com/icrowley/fake"
	"os"
	"strings"
)

const (
	GeneratorText   = "****Persona Generator****"
	GeneratorDesign = "-*-*-*-*-*-*-*-*-*-*-*-*-"
	Inputtext       = "Please enter a number:"
)

var (
	minNum  = 1
	maxNum  = 10
	countc  = 0
	ukc     = 0
	us0     = 0
	malec   = 0
	femalec = 0
)

func main() {

	r := returnnum()
	fmt.Println(GeneratorText)
	for i := 0; i < r; i++ {
		fmt.Println(persona(), "\n")
		fmt.Println("EOF Record Number:", i+1)
		fmt.Println(GeneratorDesign, "\n")
	}
	fmt.Printf("Total records successfully printed %v.\n Count of males were %v.\n Count of females were %v.\n", r, malec, femalec)
	fmt.Printf("Total count of people who reside in the US and UK are %v.\n Number based in the US are %v.\n Number based in the UK are %v.\n", countc, us0, ukc)
}

func persona() string {

	fullname := strings.ToUpper(fake.FullNameWithPrefix())
	d := strings.HasPrefix(fullname, "MR. DR. ")
	if d == true {
		malec++
	} else {
		femalec++
	}

	coadd := fake.Country()
	if coadd == "United States of America" || coadd == "United Kingdom" {
		countc++
	}
	if coadd == "United States of America" {
		us0++
	}
	if coadd == "United Kingdom" {
		ukc++
	}

	email := fake.EmailAddress()
	sadd := fake.StreetAddress()
	padd := fake.Zip()
	cadd := fake.State()
	honum := fake.Phone()
	monum := fake.DigitsN(11)

	return "\n" + "NAME: " + fullname + "\n" + "Email Address: " +
		email + "\n" + "Street Address: " + sadd + "\n" + "Zip: " + padd + "\n" + "State: " +
		cadd + "\n" + "Country: " + coadd + "\n" + "H Phone: " + honum + "\n" + "Mobile: " + monum
}

func returnnum() int {
	var num int
	fmt.Print(Inputtext)
	n, err := fmt.Scan(&num)
	if err != nil || n != 1 || n < minNum || n > maxNum {
		fmt.Printf("Invalid number, Enter between %d and %d\n", minNum, maxNum)
		os.Exit(1)
	}
	return num
}
