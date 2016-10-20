package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	p1 := Player{"Tom", 0}
	p2 := Player{"Alex", 0}

	p1PointsChannel := make(chan int)
	p2PointsChannel := make(chan int)

	var p1TotalScore int
	var p2TotalScore int

	go play(p1, p1PointsChannel)
	go play(p2, p2PointsChannel)

	for {
		p1Score := <-p1PointsChannel
		p1TotalScore = p1TotalScore + p1Score
		fmt.Printf("incremented p1Score, it's now %d \n", p1TotalScore)
		p2Score := <-p2PointsChannel
		p2TotalScore = p2TotalScore + p2Score
		fmt.Printf("incremented p2Score, it's now %d \n", p2TotalScore)
	}

	time.Sleep(30 * time.Second)
	fmt.Println("Game is over")
}

func play(player Player, pointsChannel chan int) {
	fmt.Printf("%s is ready to play\n", player.Name)
	for {
		pointsChannel <- rollTheDie(player)
	}
}

func rollTheDie(player Player) int {
	randomNumber := rand.Intn(6) + 1

	secondsToSleep := 6 - randomNumber
	fmt.Printf("%s (%d) rolled a %d, waiting %d sec\n", player.Name, player.Score, randomNumber, secondsToSleep)
	time.Sleep(time.Duration(secondsToSleep) * time.Second)
	return randomNumber
}

type Player struct {
	Name  string
	Score int
}
