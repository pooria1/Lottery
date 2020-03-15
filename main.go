package main

import (
	"fmt"
	"math/rand"
)

//addBall func will add n balls with c color to the collection

func addBall(b map[string]int) {

	fmt.Println("enter color of balls you want add: ")
	var c string
	var n int
	fmt.Print("ball: ")
	fmt.Scanf("%s", &c)
	fmt.Print("number: ")
	fmt.Scanf("%d", &n)
	b[c] += n

}

//removeBall func will remove n balls with c color from the collection.
//this func returns a boolean. It is true if balls successfully removed and false if not removed.

func removeBall(b map[string]int) bool {
	fmt.Println("enter color of balls you want remove: ")
	var c string
	var n int
	fmt.Print("ball: ")
	fmt.Scanf("%s", &c)
	fmt.Print("number: ")
	fmt.Scanf("%d", &n)

	if b[c] >= n {
		b[c] -= n
		if b[c] == 0 {
			delete(b, c)
		}
		return true
	}

	fmt.Println("can't remove")
	return false
}

//getRandBall func will return color of random ball. This func will calculate a number
//between 1 and number of balls then will find which ball is that then return its color.

func getRandBall(b map[string]int) string {

	totalBalls := 0
	for _, v := range b {
		totalBalls += v
	}

	randBallNum := rand.Intn(totalBalls) + 1
	var randBall string

	for key, value := range b {
		if randBallNum <= value {
			randBall = key
			break
		}
		randBallNum -= value
	}
	fmt.Println("Random ball is", randBall, "with", b[randBall], "points.")
	b[randBall]--
	if b[randBall] == 0 {
		delete(b, randBall)
	}
	return randBall
}

func main() {

	b := make(map[string]int)

	for {
		fmt.Println("\n[1] Add ball\n[2] Remove ball\n[3] get random ball\n[other] quit")
		var choice int
		fmt.Scanf("%d", &choice)

		switch choice {
		case 1:
			addBall(b)
		case 2:
			removeBall(b)
		case 3:
			getRandBall(b)
		default:
			return
		}
	}
}
