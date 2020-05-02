package main

// CurrentGame - singleton
var CurrentGame Game

// Game Type which represents game
type Game struct {
	currentPlayer Player
	locations []Location
	cli *CLI
}

// StartNewGame Game factory function
func StartNewGame() {
	CurrentGame.cli = NewCLI()
	// CurrentGame.currentPlayer = *NewPlayer(CurrentGame.cli.GetUserNickname())
}
