package main

import (
	"fmt"
	"github.com/icrowley/fake"
	"strings"
)

func main() {

	var num int
	fmt.Print("Please enter a number: ")
	fmt.Scan(&num)

	for i := 0; i < num; i++ {
		fmt.Println("****Persona Generator****", persona())
		fmt.Println("------------------------")
	}
}

func persona() string {

	fullname := strings.ToUpper(fake.MaleFullNameWithPrefix())
	email := fake.EmailAddress()
	sadd := fake.StreetAddress()
	padd := fake.Zip()
	cadd := fake.State()
	coadd := fake.Country()
	honum := fake.Phone()
	monum := fake.DigitsN(11)

	return "\n" + fullname + "\n" + email +
		"\n" + sadd + "\n" + padd + "\n" + cadd + "\n" + coadd +
		"\n" + honum + "\n" + monum
}

//err := fake.SetLang("ru")
//if err != nil {
//panic(err)
//}
