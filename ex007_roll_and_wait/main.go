package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	go play("Tom")
	go play("Chris")
	time.Sleep(30 * time.Second)
	fmt.Println("Game is over")
}

func play(playerName string) {
	fmt.Println(playerName + " is ready to play")
	totalScore := 0
	for {
		totalScore = rollTheDie(totalScore, playerName)
	}
}

func rollTheDie(totalScore int, playerName string) int {
	randomNumber := rand.Intn(6) + 1
	fmt.Println(playerName + " (" + strconv.Itoa(totalScore) + ")  rolled a " + strconv.Itoa(randomNumber))
	time.Sleep(time.Duration(6-randomNumber) * time.Second)
	return totalScore + randomNumber
}
