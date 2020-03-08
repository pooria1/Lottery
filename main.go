package main

import (
	"fmt"
	"math/rand"
)

type balls map[string]int

//addBall func will add n balls with c color to the collection

func addBall(b balls) {

	fmt.Println("enter color of balls you want add: ")
	var c string
	var n int
	fmt.Print("ball: ")
	fmt.Scanf("%s", &c)
	fmt.Print("number: ")
	fmt.Scanf("%d", &n)
	b[c] += n

}

//addBall func will remove n balls with c color from the collection.
//this func returns a boolean. It is true if balls successfully removed and false if not removed.

func removeBall(b balls) bool {
	fmt.Println("enter color of balls you want remove: ")
	var c string
	var n int
	fmt.Print("ball: ")
	fmt.Scanf("%s", &c)
	fmt.Print("number: ")
	fmt.Scanf("%d", &n)
	isRemoved := false
	for key, value := range b {
		if key == c && value >= n && n >= 0 {
			b[c] -= n
			if b[c] == 0 {
				delete(b, c)
			}
			isRemoved = true
			break
		}
	}
	if isRemoved == false {
		fmt.Println("can't remove")
	}
	return isRemoved
}

//getRandBall func will return color of random ball. This func will calculate a number
//between 1 and number of balls then will find which ball is that then return its color.

func getRandBall(b balls) {

	totalBalls := 0
	for _, v := range b {
		totalBalls += v
	}

	randBallNum := rand.Intn(totalBalls) + 1
	var randBall string

	for key, value := range b {
		if randBallNum <= value {
			randBall = key
		}
		randBallNum -= value
	}
	fmt.Println("Random ball is", randBall, "with", b[randBall], "points.")
	b[randBall]--
}

func main() {

	b := make(map[string]int)

	for {
		fmt.Println("[1] Add ball\n[2] Remove ball\n[other] quit")
		var choice int
		fmt.Scanf("%d", &choice)

		if choice == 1 {

			addBall(b)
			getRandBall(b)

		} else if choice == 2 {
			isRemoved := removeBall(b)
			if isRemoved {
				getRandBall(b)
			}
		} else {
			return
		}
	}
}
