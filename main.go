package main

import (
	"fmt"
	"math/rand"
)

type balls []string

//getNumBalls func calculate point of a ball with "color" color

func (b *balls) getNumBalls(color string) int {
	ans := 0
	for _, v := range *b {
		if v == color {
			ans++
		}
	}
	return ans
}

func main() {

	var b balls

	for {
		fmt.Println("[1] Add ball\n[2] quit")
		var choice int
		fmt.Scanf("%d", &choice)

		if choice == 1 {

			//here we will get the balls and add them in collection

			fmt.Println("enter color of balls you want add(type 0 to finish adding and get a random ball): ")
			var c string
			for i := 1; ; i++ {
				fmt.Printf("ball %d: ", i)
				fmt.Scanf("%s", &c)
				if c == "0" {
					break
				}
				b = append(b, c)
			}

			//now we will return a random ball with its point and remove that from collection

			i := rand.Intn(len(b))
			fmt.Println("Random ball is", b[i], "with", b.getNumBalls(b[i]), "points.")
			b = append(b[0:i], b[i+1:len(b)]...)

		} else {
			return
		}
	}
}
