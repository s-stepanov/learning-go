package main

import (
	"fmt"
	"time"
)

const LocationsCount = 5

func initGame() {
	CurrentGame = NewGame()

	kitchen := NewLocation("кухня")
	room := NewLocation("комната")
	outside := NewLocation("улица")
	hallway := NewLocation("коридор")
	home := NewLocation("домой")

	kitchen.SetDescription("ты находишься на кухне, на столе чай, надо собрать рюкзак и идти в универ.")
	kitchen.SetWelcomeMessage("кухня, ничего интересного.")
	kitchen.SetEmptyMessage("ты находишься на кухне, на столе чай, надо идти в универ.")
	hallway.SetWelcomeMessage("ничего интересного.")
	room.SetWelcomeMessage("ты в своей комнате.")
	room.SetEmptyMessage("пустая комната.")
	outside.SetWelcomeMessage("на улице весна.")

	kitchen.AddNearbyLocation(hallway)
	hallway.AddNearbyLocation(kitchen)
	hallway.AddNearbyLocation(room)
	hallway.AddNearbyLocation(outside)
	hallway.AddLock("дверь", outside)
	room.AddNearbyLocation(hallway)
	outside.AddNearbyLocation(home)

	room.AddItem("на столе", "ключи", false)
	room.AddItem("на столе", "конспекты", false)
	room.AddItem("на стуле", "рюкзак", true)

	CurrentGame.locations = append(CurrentGame.locations, *kitchen, *room, *hallway, *outside, *home)
}

func addPlayer(player *Player) {
	CurrentGame.players[player.GetNickname()] = player
	CurrentGame.locations[0].addPlayerToLocation(player)
	CurrentGame.players[player.GetNickname()].SetInitialLocation(&CurrentGame.locations[0])
}

func main() {
	initGame()
	player := NewPlayer("Tristan")
	addPlayer(player)

	go func() {
		res := player.GetOutput()
		for i := range res {
			fmt.Println(i)
		}
	}()

	player.HandleInput("идти коридор")

	time.Sleep(time.Millisecond)
}
