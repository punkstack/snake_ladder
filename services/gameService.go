package services

import (
	"errors"
	"log"
	"math/rand"
	"snakeLadder/models"
	"time"
)

type GameServiceImpl struct{}

func NewGameService() *GameServiceImpl {
	return &GameServiceImpl{}
}

func (s *GameServiceImpl) CreateGame(config models.GameConfig) (*models.Game, error) {
	dices := config.DiceConfig

	game := &models.Game{
		BoardSize:  config.BoardSize,
		Dices:      dices,
		MaxPlayers: config.MaxPlayers,
		JumpConfig: struct {
			Ladder []models.Jump
			Snake  []models.Jump
		}(config.JumpConfig),
	}

	return game, nil
}

func (s *GameServiceImpl) AddPlayer(game *models.Game, player *models.Player) error {
	if len(game.Players) >= game.MaxPlayers {
		return errors.New("maximum players reached")
	}
	game.Players = append(game.Players, *player)
	return nil
}

func (s *GameServiceImpl) StartGame(game *models.Game) string {
	rand.Seed(time.Now().UnixNano())
	for {
		for i := range game.Players {
			dice := s.getPlayerDice(game, &game.Players[i])
			roll := s.rollDice(dice)
			log.Printf("Player %s rolled a %d", game.Players[i].Name, roll)

			if game.Players[i].Position+roll > game.BoardSize {
				log.Printf("Player %s's roll of %d exceeds the board size. Ignoring roll.", game.Players[i].Name, roll)
				continue
			}

			game.Players[i].Position += roll
			log.Printf("Player %s moved to position %d", game.Players[i].Name, game.Players[i].Position)

			s.checkForJumps(game, &game.Players[i])

			if s.checkForWin(game, &game.Players[i]) {
				log.Printf("Player %s has won the game!", game.Players[i].Name)
				return game.Players[i].Name
			}
		}
	}
}

func (s *GameServiceImpl) getPlayerDice(game *models.Game, player *models.Player) models.Dice {
	for _, d := range game.Dices {
		if d.Type == player.DiceEquiped {
			return d
		}
	}
	return models.Dice{}
}

func (s *GameServiceImpl) rollDice(dice models.Dice) int {
	return dice.AvailableRolls[rand.Intn(len(dice.AvailableRolls))]
}

func (s *GameServiceImpl) checkForJumps(game *models.Game, player *models.Player) {
	for _, ladder := range game.JumpConfig.Ladder {
		if player.Position == ladder.Start {
			player.Position = ladder.End
			log.Printf("Player %s climbed a ladder to position %d", player.Name, player.Position)
		}
	}
	for _, snake := range game.JumpConfig.Snake {
		if player.Position == snake.Start {
			player.Position = snake.End
			log.Printf("Player %s slid down a snake to position %d", player.Name, player.Position)
		}
	}
}

func (s *GameServiceImpl) checkForWin(game *models.Game, player *models.Player) bool {
	return player.Position == game.BoardSize
}
