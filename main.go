package main

import (
	"math/rand"
	"time"
	"fmt"
)

type PlayerValues struct {
	Position  int
	Gold      int
	Ownership string
}

// ToDo: Improve print messages
func main() {

	var player1 = PlayerValues{0,300,""}
	var player2 = PlayerValues{0,300,""}
	var player3 = PlayerValues{0,300,""}
	var player4 = PlayerValues{0,300,""}
	player1.Position = 0
	rand.Seed(time.Now().UTC().UnixNano())

	var counter = 0

	var players []PlayerValues

	players = append(players, player1)
	players = append(players, player2)
	players = append(players, player3)
	players = append(players, player4)

	for !gameIsOver() {
		fmt.Println("Player ",counter+1, " is playing")
		doPlayerTurn(&players[counter])
		counter = (counter + 1) % len(players)
	}
}

// ToDo: Implement. Stop game when only one player has money left
func gameIsOver() bool {
	return false
}

func doPlayerTurn(player *PlayerValues)  {
	_, err := fmt.Scanln()
	if err != nil {
		fmt.Println("There was an error!")
	}
	dice1 := rollTheDice()
	dice2 := rollTheDice()
	player.Position = player.Position + dice1 + dice2
	// ToDo: Replace number by streetname
	// ToDo: Change action if street already purchased
	fmt.Println(player.Position, "is free to buy", "Press 1 to buy or move")
	var input string
	fmt.Scanln(&input)
	if input =="1"{
		// ToDo: Grant ownership on purchase
		player.Gold=player.Gold-300
	}
	fmt.Println("You got",player.Gold)

	for dice1 == dice2 {
		dice1 = rollTheDice()
		dice2 = rollTheDice()
		player.Position = player.Position + dice1 + dice2
	}
}

func rollTheDice() int {
	var diceRoll int
	diceRoll = rand.Int()%6 + 1
	fmt.Println(diceRoll)
	return diceRoll
}
