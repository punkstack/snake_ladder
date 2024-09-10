package models

type DiceType string

const (
	CROOKED_DICE DiceType = "crooked"
	NORMAL_DICE  DiceType = "normal"
)

type Dice struct {
	Type           DiceType `json:"type"`
	AvailableRolls []int    `json:"availableRolls"`
}
