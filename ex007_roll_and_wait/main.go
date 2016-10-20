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
	go play(p1)
	go play(p2)
	time.Sleep(30 * time.Second)
	fmt.Println("Game is over")
}

func play(player Player) {
	fmt.Println(player.Name + " is ready to play")
	for {
		player.Score = rollTheDie(player)
	}
}

func rollTheDie(player Player) int {
	randomNumber := rand.Intn(6) + 1
	secondsToSleep := 6 - randomNumber
	fmt.Printf("%s (%d) rolled a %d, waiting %d sec\n", player.Name, player.Score, randomNumber, secondsToSleep)
	time.Sleep(time.Duration(secondsToSleep) * time.Second)
	return player.Score + randomNumber
}

type Player struct {
	Name  string
	Score int
}
