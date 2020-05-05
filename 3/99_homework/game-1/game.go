package main

// CurrentGame - singleton
var CurrentGame *Game


// Game Type which represents game
type Game struct {
	players map[string]*Player
	locations []Location
}

func NewGame() *Game {
	game := new(Game)
	game.players = make(map[string]*Player)

	return game
}
