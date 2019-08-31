package main

import (
	"fmt"
	"math/rand"
	"time"
)

type PlayerValues struct {
	Position  int
	Gold      int
	Ownership string
}

type Street struct {
	Name  string
	Price int
	Type  string
	Color string
}

var players []PlayerValues

// ToDo: Improve print messages
func main() {

	var player1 = PlayerValues{0, 300, ""}
	var player2 = PlayerValues{0, 300, ""}
	var player3 = PlayerValues{0, 300, ""}
	var player4 = PlayerValues{0, 300, ""}
	player1.Position = 0
	rand.Seed(time.Now().UTC().UnixNano())

	var counter = 0

	players = append(players, player1)
	players = append(players, player2)
	players = append(players, player3)
	players = append(players, player4)

	for !gameIsOver() {
		fmt.Println("Player ", counter+1, " is playing")
		doPlayerTurn(&players[counter])
		counter = (counter + 1) % len(players)
	}
	fmt.Println("Das Spiel ist vorbei.")
	for index := range players {
		if players[index].Gold > 0 {
			fmt.Println("Player with number ", index+1, "has won. He has ", players[index].Gold, " money at the end of the game")
		}
	}
	fmt.Println("Thanks for playing")
}

// ToDo: Implement. Stop game when only one player has money left
func gameIsOver() bool {
	var numberOfPlayersWithMoney = 0
	for index := range players {
		if players[index].Gold > 0 {
			numberOfPlayersWithMoney = numberOfPlayersWithMoney + 1
		}
	}
	return numberOfPlayersWithMoney <= 1
}

func doPlayerTurn(player *PlayerValues) {
	if player.Gold > 0 {
		fmt.Println("You are on field ", player.Position, " and have ", player.Gold, " money. Press enter to roll the dice.")
		_, err := fmt.Scanln()
		if err != nil {
			fmt.Println("There was an error!")
		}
		dice1 := rollTheDice()
		dice2 := rollTheDice()
		movePlayer(player, dice1, dice2)
		// ToDo: Change action if street already purchased
		// Pick a card on card fields - dont allow purchase
		// pay taxes on tax fields
		// Free parking, get money from the bank
		// Get money on LOS
		doSomethingWithStreets(player)

		for dice1 == dice2 {
			dice1 = rollTheDice()
			dice2 = rollTheDice()
			movePlayer(player, dice1, dice2)
			doSomethingWithStreets(player)
		}
	} else {
		fmt.Println(" This player is broke, skipping")
	}
}

// ToDo: Rename function, implement all field types
func doSomethingWithStreets(player *PlayerValues) {
	street := getStreetByPosition(player.Position)
	if street.Type == "Strasse" {
		fmt.Println(street.Name, "is free to buy", "Press 1 to buy or move")
		var input string
		fmt.Scanln(&input)
		if input == "1" {
			// ToDo: Grant ownership on purchase
			player.Gold = player.Gold - street.Price
		}

	}
	// ToDo: Implement more social cards
	if street.Type == "Gemeinschaftsfeld" {
		fmt.Println("Its your Birthday you get 200 from every player")
		for index, _ := range players {
			players[index].Gold = players[index].Gold - 200

		}
		player.Gold = player.Gold + len(players)*200

	}
	fmt.Println("You got", player.Gold)
}

func movePlayer(player *PlayerValues, dice1 int, dice2 int) {
	player.Position = (player.Position + dice1 + dice2) % 40
}

func getStreetByPosition(position int) Street {
	var streetDictonary []Street
	streetDictonary = append(streetDictonary, Street{"Los", 0, "Los", "None"})
	streetDictonary = append(streetDictonary, Street{"Badstrasse", 60, "Strasse", "Lila"})
	streetDictonary = append(streetDictonary, Street{"Gemeinschaftsfeld", 0, "Gemeinschaftsfeld", "None"})
	streetDictonary = append(streetDictonary, Street{"Turmstrasse", 60, "Strasse", "Lila"})
	streetDictonary = append(streetDictonary, Street{"Einkomensteuer", 200, "Strafe", "None"})

	streetDictonary = append(streetDictonary, Street{"Südbahnhof", 200, "Bahnhof", "None"})
	streetDictonary = append(streetDictonary, Street{"Chauseestrasse", 100, "Strasse", "Hellblau"})
	streetDictonary = append(streetDictonary, Street{"Ereignissfeld", 0, "Ereignissfeld", "None"})
	streetDictonary = append(streetDictonary, Street{"Elisenstrasse", 100, "Strasse", "Hellblau"})
	streetDictonary = append(streetDictonary, Street{"Poststrasse", 120, "Strasse", "Hellblau"})

	streetDictonary = append(streetDictonary, Street{"Gefängnis", 0, "Gefängnis", "None"})
	streetDictonary = append(streetDictonary, Street{"Seestrasse", 140, "Strasse", "Pink"})
	streetDictonary = append(streetDictonary, Street{"Elektizitätswerk", 150, "Werk", "None"})
	streetDictonary = append(streetDictonary, Street{"Hafenstrasse", 140, "Strasse", "Pink"})
	streetDictonary = append(streetDictonary, Street{"Neue Strasse", 160, "Strasse", "Pink"})

	streetDictonary = append(streetDictonary, Street{"Westbahnhof", 200, "Bahnhof", "None"})
	streetDictonary = append(streetDictonary, Street{"Münchenerstrasse", 180, "Strasse", "Orange"})
	streetDictonary = append(streetDictonary, Street{"Gemeinschaftsfeld", 0, "Gemeinschaftsfeld", "None"})
	streetDictonary = append(streetDictonary, Street{"Wienerstrasse", 180, "Strasse", "Orange"})
	streetDictonary = append(streetDictonary, Street{"Berlinerstrasse", 200, "Strasse", "Orange"})

	streetDictonary = append(streetDictonary, Street{"Frei Parken", 0, "Frei Parken", "None"})
	streetDictonary = append(streetDictonary, Street{"Theaterstrasse", 220, "Strasse", "Rot"})
	streetDictonary = append(streetDictonary, Street{"Ereignissfeld", 0, "Ereignissfeld", "None"})
	streetDictonary = append(streetDictonary, Street{"Museumstrasse", 220, "Strasse", "Rot"})
	streetDictonary = append(streetDictonary, Street{"Opernplatz", 240, "Strasse", "Rot"})

	streetDictonary = append(streetDictonary, Street{"Nordbahnhof", 200, "Bahnhof", "None"})
	streetDictonary = append(streetDictonary, Street{"Lessingstrasse", 260, "Strasse", "Gelb"})
	streetDictonary = append(streetDictonary, Street{"Schillerstrasse", 260, "Strasse", "Gelb"})
	streetDictonary = append(streetDictonary, Street{"Wasserwerk", 150, "Werk", "None"})
	streetDictonary = append(streetDictonary, Street{"Ghoetestrasse", 280, "Strasse", "Gelb"})

	streetDictonary = append(streetDictonary, Street{"Gehe ins Gefängnis", 0, "Gehe ins Gefängnis", "None"})
	streetDictonary = append(streetDictonary, Street{"Rathausplatz", 300, "Strasse", "Grün"})
	streetDictonary = append(streetDictonary, Street{"Hauptstrasse", 300, "Strasse", "Grün"})
	streetDictonary = append(streetDictonary, Street{"Gemeinschaftsfeld", 0, "Gemeinschaftsfeld", "None"})
	streetDictonary = append(streetDictonary, Street{"Bahnhofstrasse", 320, "Strasse", "Grün"})

	streetDictonary = append(streetDictonary, Street{"Hauptbahnhof", 200, "Bahnhof", "None"})
	streetDictonary = append(streetDictonary, Street{"Ereignissfeld", 0, "Ereignissfeld", "None"})
	streetDictonary = append(streetDictonary, Street{"Parkstrasse", 350, "Strasse", "Dunkelblau"})
	streetDictonary = append(streetDictonary, Street{"Zusatzsteuer", 100, "Strafe", "None"})
	streetDictonary = append(streetDictonary, Street{"Schlossallee", 400, "Strasse", "Dunkelblau"})

	return streetDictonary[position%len(streetDictonary)]
}

func rollTheDice() int {
	var diceRoll int
	diceRoll = rand.Int()%6 + 1
	fmt.Println(diceRoll)
	return diceRoll
}
