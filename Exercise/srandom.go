package main

import "github.com/icrowley/fake"

type persona1 struct {
	fullname      string
	email         string
	StreetAddress string
	Zip           string
	State         string
	Country       string
	Phone         int
	Mobile        int
}

func (m *persona1) Getfullname() string {
	if m != nil {
		return fake.FullNameWithPrefix()
	}
	return ""
}
func (m *persona1) Getemail() string {
	if m != nil {
		return fake.EmailAddress()
	}
	return ""
}
func (m *persona1) GetStreetAddress() string {
	if m != nil {
		return fake.StreetAddress()
	}
	return ""
}
func (m *persona1) GetZip() string {
	if m != nil {
		return fake.Zip()
	}
	return ""
}
func (m *persona1) GetState() string {
	if m != nil {
		return fake.State()
	}
	return ""
}
func (m *persona1) GetCountry() string {
	if m != nil {
		return fake.Country()
	}
	return ""
}
func (m *persona1) GetPhone() string {
	if m != nil {
		return fake.Phone()
	}
	return ""
}
func (m *persona1) GetMobile() string {
	if m != nil {
		return fake.DigitsN(11)
	}
	return ""
}
