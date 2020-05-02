package main

import (
	"fmt"
)

type Player struct {
	nickname string
	currentLocation *Location
	inventory []*Item
	hasBackpack bool
}

func NewPlayer(nickname string, initialLocation *Location) *Player {
	player := new(Player)
	player.nickname = nickname
	player.currentLocation = initialLocation

	return player
}

func (player *Player) SwitchLocation(locationName string) string {
	if !(player.currentLocation.HasNearbyLocation(locationName)) {
		return fmt.Sprintf("нет пути в %s", locationName)
	}
	
	for _, location := range player.currentLocation.GetNearbyLocations() {
		if (location.GetName() == locationName) {
			player.currentLocation = location
			return player.currentLocation.GetDescription() + player.currentLocation.GetNearbyLocationsString()
		}
	}

	return fmt.Sprintf("нет пути в %s", locationName)
}

func (player *Player) LookAround() (result string) {
	result = player.currentLocation.GetAvailableItemsString()

	result += player.currentLocation.GetNearbyLocationsString()

	return result
}

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
	default:
		return "неизвестная команда"
	}
}
