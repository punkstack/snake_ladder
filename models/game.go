package models

type GameConfig struct {
	BoardSize  int    `json:"boardSize"`
	DiceConfig []Dice `json:"diceConfig"`
	MaxPlayers int    `json:"maxPlayers"`
	JumpConfig struct {
		Ladder []Jump `json:"ladder"`
		Snake  []Jump `json:"snake"`
	} `json:"jumpConfig"`
}

type Game struct {
	BoardSize  int
	Dices      []Dice
	MaxPlayers int
	Players    []Player
	JumpConfig struct {
		Ladder []Jump
		Snake  []Jump
	}
}
