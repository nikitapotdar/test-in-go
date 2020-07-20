package main

import "fmt"
		type deck []string

		func newDeck() deck{
			cards := deck{}
			cardSuits :=[]string{"spades","dimonds","hrart","cubes"}
			cardValue :=[]string{"Ace","two","three","four"}
			for _,suit := range cardSuits {
				for _,value :=range cardValue{
					cards=append(cards,value + "of" + suit)
				}
				
			}
			return cards
		}
		func (d deck) print(){
			for i,card :=range d{
				fmt.Println(i,card)
			}
		}
		
