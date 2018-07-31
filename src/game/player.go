package game

type Player struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

func NewPlayer(name, symbol string) Player {
	return Player{
		Name:   name,
		Symbol: symbol,
	}
}
