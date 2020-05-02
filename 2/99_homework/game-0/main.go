package main

import (
	"strings"
)

func initGame() {
	CurrentGame = *new(Game)

	kitchen := NewLocation("кухня")
	room := NewLocation("комната")
	outside := NewLocation("улица")
	hallway := NewLocation("коридор")
	home := NewLocation("домой")

	kitchen.SetDescription("ты находишься на кухне, на столе чай, надо собрать рюкзак и идти в универ.")
	hallway.SetDescription("ничего интересного.")
	room.SetDescription("ты в своей комнате.")
	outside.SetDescription("на улице весна.")

	kitchen.AddNearbyLocation(hallway)
	hallway.AddNearbyLocation(kitchen)
	hallway.AddNearbyLocation(room)
	hallway.AddNearbyLocation(outside)
	room.AddNearbyLocation(hallway)
	outside.AddNearbyLocation(home)

	room.AddItem("на столе", "ключи", false)
	room.AddItem("на столе", "конспекты", false)
	room.AddItem("на стуле", "рюкзак", true)

	CurrentGame.currentPlayer = *NewPlayer("player", kitchen)

	CurrentGame.locations = append(CurrentGame.locations, *kitchen, *room, *hallway, *outside, *home)
}

func handleCommand(command string) (result string) {
	action := strings.Split(command, " ")
	return CurrentGame.currentPlayer.PerformAction(action[0], action[1:]...)
}

func main() {
	initGame()
}