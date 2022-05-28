package main

import "fmt"



func main(){
	cards := []string{"Ace of Hearts",newCard()}

	for i,card := range cards{
		fmt.Println(i,card)
	}
	
}

func newCard() string{
	return "Five of Diamonds"
}