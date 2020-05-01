package main

// CurrentGame - singleton
var CurrentGame Game

// Game Type which represents game
type Game struct {
	currentPlayer *Player
	cli *CLI
}

// StartNewGame Game factory function
func StartNewGame() {
	CurrentGame := new(Game)
	CurrentGame.cli = NewCLI()
	CurrentGame.currentPlayer = NewPlayer(CurrentGame.cli.GetUserNickname())
}

// Player Type which represents player
type Player struct {
	nickname string
}

// NewPlayer Player object constructor
func NewPlayer(nickname string) *Player {
	player := new(Player)
	player.nickname = nickname

	return player
}