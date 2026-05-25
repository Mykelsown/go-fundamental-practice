package main

import "fmt"

type address struct {
	country string
	state   string
	LGA     string
	town    string
}


func (f address) PrintDetails() {
	fmt.Println(f.country)
	fmt.Println(f.state)
	fmt.Println(f.LGA)
	fmt.Println(f.town)
}

type employee struct {
	name       string
	department string
	age        int
	fAddress   address
	cAddress   address
}


func main() {
	// composition
	empDetails := employee{
		name:       "Micheal Samuel",
		department: "Engineering",
		age:        14,
		fAddress:  address{
			country: "Nigeria",
			state: "Osun",
			LGA: "Bolorun duro",
			town: "block 1",
		} ,
		cAddress: address{
			country: "Nigeria",
			state: "Lagos",
			LGA: "Yaba",
			town: "Montgometry",
		},
	}
	empDetails.cAddress.PrintDetails()
	empDetails.fAddress.PrintDetails()
}