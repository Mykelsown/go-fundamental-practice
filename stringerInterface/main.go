package main

import "fmt"

type Game interface {
	GetData()
}

type ShootingGame struct {
	Name string
	Rating string
	Downloads string
}

func (sg ShootingGame) GetData() {
	fmt.Println(sg.Name)
	fmt.Println(sg.Rating)
	fmt.Println(sg.Downloads)
}

func (sg ShootingGame) String() string{
	return fmt.Sprintf("Name: %s, Rating: %s, Downloads: %s", sg.Name, sg.Rating, sg.Downloads)
}

func main() {
	shGm := ShootingGame{
		Name: "Call of Duty",
		Rating: "4.3/5",
		Downloads: "19 million",
	}

	fmt.Println(shGm)
	shGm.GetData()
}
