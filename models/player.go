package models

type Player struct {
	Id          int
	Name        string
	DiceEquiped DiceType
	Position    int
}

func NewPlayer(id int, name string, diceType DiceType) (error, *Player) {
	return nil, &Player{
		Id:          id,
		Name:        name,
		DiceEquiped: diceType,
		Position:    0,
	}
}
