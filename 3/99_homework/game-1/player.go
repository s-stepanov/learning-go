package main

import (
	"fmt"
	"strings"
)

// Player struct
type Player struct {
	nickname string
	currentLocation *Location
	inventory []*Item
	hasBackpack bool
	outputChannel chan string
}

// NewPlayer player constructor
func NewPlayer(nickname string) *Player {
	player := new(Player)
	player.nickname = nickname
	player.hasBackpack = false
	player.outputChannel = make(chan string)

	return player
}

// GetNickname returns player's nickname
func (player *Player) GetNickname() string {
	return player.nickname
}

// SetInitialLocation sets initial location
func (player *Player) SetInitialLocation(location *Location) {
	player.currentLocation = location
}

// SwitchLocation If there is a way from current location to another, moves player to the passed one
func (player *Player) SwitchLocation(locationName string) string {
	locationToGo := player.currentLocation.GetNearbyLocation(locationName)

	if locationToGo == nil {
		return fmt.Sprintf("нет пути в %s", locationName)
	}

	for _, lock := range player.currentLocation.locks {
		if lock.locationToLock.GetName() == locationName {
			return "дверь закрыта"
		}
	}

	player.currentLocation.removePlayerFromLocation(player)
	locationToGo.addPlayerToLocation(player)

	player.currentLocation = locationToGo
	return player.currentLocation.GetWelcomeMessage() + player.currentLocation.GetNearbyLocationsString()
}

// LookAround look around in current location
func (player *Player) LookAround() (result string) {
	result = player.currentLocation.GetAvailableItemsString()

	if result == "" && !(player.currentLocation.GetName() == "кухня" && !player.hasBackpack) {
		return player.currentLocation.GetEmptyMessage() + player.currentLocation.GetNearbyLocationsString()
	}

	result += player.currentLocation.GetDescription() + player.currentLocation.GetNearbyLocationsString()

	var playersString string
	var filteredPlayersInLocation []*Player

	for _, p := range player.currentLocation.getPlayersInLocation() {
		if strings.Compare(player.GetNickname(), p.GetNickname()) != 0 {
			filteredPlayersInLocation = append(filteredPlayersInLocation, p)
		}
	}

	if len(filteredPlayersInLocation) > 0 {
		playersString = ". Кроме вас тут ещё "
		for index, neighbor := range filteredPlayersInLocation {
			playersString += neighbor.GetNickname()
			if index != len(filteredPlayersInLocation) - 1 {
				playersString += ", "
			}
		}
	}

	return result + playersString
}

// PickupItem picks up generic item
func (player *Player) PickupItem(itemName string) (result string) {
	location, index, item := player.currentLocation.FindItem(itemName) 

	if (item == nil) {
		return "нет такого"
	}

	if (!player.hasBackpack) {
		return "некуда класть"
	}

	items := player.currentLocation.GetAvailableItems()[location]
	player.currentLocation.GetAvailableItems()[location] = append(items[:index], items[index+1:]...)
	
	if len(player.currentLocation.GetAvailableItems()[location]) == 0 {
		delete(player.currentLocation.GetAvailableItems(), location)
	}

	player.inventory = append(player.inventory, item)

	return "предмет добавлен в инвентарь: " + item.GetName()
}

// WearItem picks up a clothing item (e.g backpack)
func (player *Player) WearItem(itemName string) (result string) {
	location, index, item := player.currentLocation.FindItem(itemName) 

	if (item == nil) {
		return "нет такого"
	}

	if !item.IsWearable() {
		return "нельзя надеть"
	}

	items := player.currentLocation.GetAvailableItems()[location]
	player.currentLocation.GetAvailableItems()[location] = append(items[:index], items[index+1:]...)

	if len(player.currentLocation.GetAvailableItems()[location]) == 0 {
		delete(player.currentLocation.GetAvailableItems(), location)
	}

	player.inventory = append(player.inventory, item)

	if itemName == "рюкзак" {
		player.hasBackpack = true
	}

	return "вы одели: " + item.GetName()
}

// UseItem unlocks the door with a key if the key is in inventory and the door is present in location
func (player *Player) UseItem(itemName string, itemToUseOn string) string {
	var inventoryItem *Item
	for _, item := range player.inventory {
		if (item.GetName() == itemName) {
			inventoryItem = item
			break
		}
	}

	if (inventoryItem == nil) {
		return "нет предмета в инвентаре - " + itemName
	}

	if itemName == "ключи" && itemToUseOn == "дверь" {
		player.currentLocation.Unlock()
		return "дверь открыта"
	}

	return "не к чему применить"
}

// SayToAll sends message to all players in current location
func (player *Player) SayToAll(messages ...string) {
	message := strings.Join(messages, " ")
	for _, receiver := range player.currentLocation.getPlayersInLocation() {
		receiver.outputChannel <- player.GetNickname() + " говорит: " + message 
	}
}

// SayToPlayer sends message to specific player in current location
func (player *Player) SayToPlayer(params ...string) {
	playerName := params[0]
	var message string
	var said bool

	if len(params) > 1 {
		message = strings.Join(params[1:], " ")
	}

	for _, reciever := range player.currentLocation.getPlayersInLocation() {
		if reciever.GetNickname() == playerName {
			said = true
			if len(message) == 0 {
				reciever.outputChannel <- player.GetNickname() + " выразительно молчит, смотря на вас"
			} else  {
				reciever.outputChannel <- player.GetNickname() + " говорит вам: " + message
			}
		}
	}
	if !said {
		player.outputChannel <- "тут нет такого игрока"
	}
}

// PerformAction calls appropriate method based on user's command
func (player *Player) PerformAction(command string, parameters ...string) (result string) {
	switch command {
	case "идти":
		return player.SwitchLocation(parameters[0])
	case "осмотреться":
		return player.LookAround()
	case "взять":
		return player.PickupItem(parameters[0])
	case "одеть":
		return player.WearItem(parameters[0])
	case "применить":
		return player.UseItem(parameters[0], parameters[1])
	default:
		return "неизвестная команда"
	}
}

// HandleInput splits incoming string into command and params, runs command in a separate goroutine
// and sends the result of command to player's output channel
func (player *Player) HandleInput(command string) {
	go func () {
		action := strings.Split(command, " ")
		if action[0] == "сказать" {
			player.SayToAll(action[1:]...)
		} else if action[0] == "сказать_игроку" {
			player.SayToPlayer(action[1:]...)
		} else {
			player.outputChannel <- player.PerformAction(action[0], action[1:]...)
		}
	}()
}

// GetOutput returns player's output channel
func (player *Player) GetOutput() <-chan string {
	return player.outputChannel
}
