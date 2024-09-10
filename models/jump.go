package models

type JumpConfig struct {
	Snake  []Jump `json:"snake"`
	Ladder []Jump `json:"ladder"`
}

type Jump struct {
	Start int `json:"start"`
	End   int `json:"end"`
}
