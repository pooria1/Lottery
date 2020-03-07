package main

import (
	"fmt"
	"math/rand"
)

type ball struct {
	color string
	ID    int
}

var ID int

type balls []ball

func (b *balls) addBall(color string) {
	ID++
	*b = append(*b, ball{color, ID})
}

func (b *balls) getNumBalls(color string) int {
	ans := 0
	for _, v := range *b {
		if v.color == color {
			ans++
		}
	}
	return ans
}

func main() {

	var b balls

	for {
		fmt.Println("[1] Add ball\n[2] get random ball\n[3] quit")
		var choice int
		fmt.Scanf("%d", &choice)

		switch choice {

		case 1:
			fmt.Print("color: ")
			var c string
			fmt.Scanf("%s", &c)
			b.addBall(c)
			break

		case 2:
			i := rand.Intn(ID)
			fmt.Println("Your ball is", b[i].color, "and point is", b.getNumBalls(b[i].color))
			break

		default:
			return
		}
	}
}
