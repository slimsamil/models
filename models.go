package models

import (
	"strconv"
)

type Stringer interface{
	String() string
}

type Address struct {
	straße string
	stadt string
}


func NewAddress(straße string, stadt string) Address{
	a := Address{
		straße,
		stadt,
	}

	return a
}

type Student struct {
	matnumber int
	name string
	address Address
	grades map[string]float32
} 

func NewStudent(matnumber int, name string, address Address, grades map[string]float32) Student{
	s := Student{
		matnumber,
		name,
		address,
		grades,
	}

	return s
}

type SharedFlat struct{
	address Address
	residents []Student
}

func NewFlat(address Address, residents []Student) SharedFlat{
	flat := SharedFlat{
		address,
		residents,
	}

	return flat
}

func (s Student) AddGrade(lecture string, grade float32) {
	s.grades[lecture] = grade
}

func (s Student) GetOverallAverageGrade() float32 {
	var average float32
	var counter float32
	for _, v := range s.grades {
		average += v
		counter++
	}

	return average / counter
}

func (s Student) String() string {
	var str string = strconv.Itoa(s.matnumber) + ", " + s.name + ", " + s.address.String()
	return str
}

func (a Address) String() string {
	return a.straße + ", " + a.stadt
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