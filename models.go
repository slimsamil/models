package models

import (
	"strconv"
)

type Stringer interface{
	String() string
}

func (s Student) String() string {
	var str string = strconv.Itoa(s.matnumber) + ", " + s.name + ", " + s.address.String()
	return str
}

func (a Address) String() string {
	return a.straße + ", " + a.stadt
}

type Address struct {
	straße string
	stadt string
}

type Student struct {
	matnumber int
	name string
	address Address
} 

type SharedFlat struct{
	address Address
	residents []Student
}



func (sf SharedFlat) AddStudent(student Student) {
	sf.residents = append(sf.residents, student)
}

func (sf SharedFlat) GetNames() []string {
	var residents []string
	for k, _ := range sf.residents {
		residents = append(residents, sf.residents[k].String())
	}
	return residents
}