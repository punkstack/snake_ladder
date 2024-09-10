package interfaces

import "snakeLadder/models"

type GameService interface {
	CreateGame(config models.GameConfig) (*models.Game, error)
	AddPlayer(game *models.Game, player *models.Player) error
	StartGame(game *models.Game) string
	rollDice(dice models.Dice) int
	checkForJumps(game *models.Game, player *models.Player)
	checkForWin(game *models.Game, player *models.Player) bool
	getPlayerDice(game *models.Game, player *models.Player) models.Dice
}
