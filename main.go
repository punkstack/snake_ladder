package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"snakeLadder/models"
	"snakeLadder/services"
)

func main() {
	// Load game configuration
	configData, err := ioutil.ReadFile("config/gameConfigOne.json")
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	var gameConfig models.GameConfig
	err = json.Unmarshal(configData, &gameConfig)
	if err != nil {
		log.Fatalf("Error unmarshalling config data: %v", err)
	}

	// New game
	gameService := services.NewGameService()
	game, err := gameService.CreateGame(gameConfig)
	if err != nil {
		log.Fatalf("Error creating game: %v", err)
	}

	err, player1 := models.NewPlayer(1, "Manoj", models.CROOKED_DICE)
	if err != nil {
		log.Fatalf("Error adding player: %v", err)
	}
	err = gameService.AddPlayer(game, player1)
	if err != nil {
		log.Fatalf("Error adding player: %v", err)
	}
	err, player2 := models.NewPlayer(2, "Suhail", models.NORMAL_DICE)
	if err != nil {
		log.Fatalf("Error adding player: %v", err)
	}
	err = gameService.AddPlayer(game, player2)
	if err != nil {
		log.Fatalf("Error adding player: %v", err)
	}

	// Start game
	winner := gameService.StartGame(game)
	fmt.Printf("The winner is: %s\n", winner)
}
