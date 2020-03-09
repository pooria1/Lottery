package main

import "testing"

func TestGetRandBall(t *testing.T) {
	b := make(map[string]int)
	b["red"] += 10
	b["blue"] += 5
	b["white"] += 2

	removedBallColor := getRandBall(b)

	x := b[removedBallColor]
	if removedBallColor == "red" {
		if x != 1 {
			t.Error("Expected", 9, "Got", x)
		}
	} else if removedBallColor == "blue" {
		if x != 4 {
			t.Error("Expected", 4, "Got", x)
		}
	} else {
		if x != 1 {
			t.Error("Expected", 1, "Got", x)
		}
	}

}

func TestRemoveBall(t *testing.T) {
	b := make(map[string]int)
	b["red"] += 10
	b["blue"] += 5
	b["white"] += 2

	removeBall(b)

	if b["red"] != 10 || b["blue"] != 5 || b["white"] != 2 {
		t.Error("unexpected removing")
	}

	b["red"] -= 2

	if b["red"] != 8 {
		t.Error("Expected", 8, "Got", b["red"])
	}
}

func TestAddBall(t *testing.T) {
	b := make(map[string]int)
	b["red"] += 10
	b["blue"] += 5
	b["white"] += 2

	b["red"] += 2
	b["yellow"] += 1500

	if b["red"] != 12 {
		t.Error("Expected", 12, "Got", b["red"])
	}

	if b["yellow"] != 1500 {
		t.Error("Expected", 1500, "Got", b["yellow"])
	}
}
