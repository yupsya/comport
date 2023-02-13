package main

import "fmt"

type cat struct{
	age int
	name string
	breed string
	vaccine bool		
}

func rooms(room string) string{
	if room == "cave" {
		fmt.Println("You are in the cave")
	} else if room == "entrance" {
		fmt.Println("You are at entrance")
	} else if room == "mountain" {
		fmt.Println("You are at the mountain")
	} else {
		fmt.Println("You can`t understand where you are")
	}
	return room
}

func main() {
	rooms("entrance")
}