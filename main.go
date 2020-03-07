package main

import (
	"fmt"
	"math/rand"
)

type balls []string

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
			fmt.Print("adding a ball with color: ")
			var c string
			fmt.Scanf("%s", &c)
			b = append(b, c)
			i := rand.Intn(len(b))
			fmt.Println("Random ball is", b[i], "with", b.getNumBalls(b[i]), "points.")
			break
		} else {
			return
		}
	}
}
