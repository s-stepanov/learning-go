package main

import (
	"fmt"
	"errors"
	"sort"
	"sync"
)

type Location struct {
	name string
	description string
	welcomeMessage string
	emptyMessage string
	nearbyLocations []*Location
	availableItems map[string][]*Item
	playersInLocation []*Player
	locks []*Lock

	mutex *sync.Mutex
}

func NewLocation(name string) *Location {
	location := new(Location)
	location.name = name
	location.description = ""
	location.availableItems = make(map[string][]*Item)

	location.mutex = &sync.Mutex{}
	
	return location
}

func (location *Location) GetName() string {
	return location.name
}

func (location *Location) SetNearbyLocations(locations []*Location) {
	location.nearbyLocations = locations
}

func (location *Location) GetNearbyLocations() []*Location {
	return location.nearbyLocations
}

// AddNearbyLocation adds new location as a nearby
func (location *Location) AddNearbyLocation(locationToAdd *Location) error {
	if location.GetNearbyLocation(locationToAdd.name) != nil  || locationToAdd.name == location.name {
		return fmt.Errorf("Cannot add nearby location with name: " + location.GetName())
	}
	location.nearbyLocations = append(location.nearbyLocations, locationToAdd)
	return nil
}

// GetNearbyLocation returns nearby location by its name if such way exists, otherwise nil
func (location *Location) GetNearbyLocation(locationName string) *Location {
	for _, location := range location.nearbyLocations {
		if location.name == locationName {
			return location
		}
	}
	return nil
}

// GetNearbyLocationsString returns nearby locations in string format
func (location *Location) GetNearbyLocationsString() string {
	result := ""
	if len(location.GetNearbyLocations()) > 0 {
		availableLocationsString := " можно пройти - "
		for index, item := range location.GetNearbyLocations() {
			availableLocationsString += item.GetName()
			if index != len(location.GetNearbyLocations()) - 1 {
				availableLocationsString += ", "
			}
		}
		result += availableLocationsString
	}
	return result
}

func (location *Location) SetDescription(description string) {
	location.description = description
}

func (location *Location) GetDescription() string {
	return location.description
}

func (location *Location) SetWelcomeMessage(message string) {
	location.welcomeMessage = message
}

func (location *Location) GetWelcomeMessage() string {
	return location.welcomeMessage
}

func (location *Location) SetEmptyMessage(message string) {
	location.emptyMessage = message
} 

func (location *Location) GetEmptyMessage() string {
	return location.emptyMessage
}

func (location *Location) AddItem(position string, itemName string, isWearable bool) {
	location.mutex.Lock()
	defer location.mutex.Unlock()

	location.availableItems[position] = append(location.availableItems[position], NewItem(itemName, isWearable))
}

func (location *Location) GetAvailableItems() map[string][]*Item {
	return location.availableItems
}

// GetAvailableItemsString returns list of items in location in string format
func (location *Location) GetAvailableItemsString() (result string) {
	mapIndex := 0
	var itemPositions []string
	for position := range location.availableItems {
		itemPositions = append(itemPositions, position)
	}
	sort.Strings(itemPositions)

	for _, itemPosition := range itemPositions {
		items := location.availableItems[itemPosition]
		if (len(items) > 0) {
			if (len(items) == 1) {
				result += itemPosition + " - "
			} else {
				result += itemPosition + ": "
			}
		}

		for index, item := range items {
			result += item.GetName()
			if index != len(items) - 1 {
				result += ", "
			}
		}
		if mapIndex != len(location.availableItems) - 1 {
			result += ", "
		}
		mapIndex++
	}

	if (len(result) != 0) {
		return result + "."
	}

	return result
}

// FindItem returns position, item and its index by name if it is present in location, otherwise nil
func (location *Location) FindItem(itemName string) (string, int, *Item) {
	for position, items := range location.GetAvailableItems() {
		for index, item := range items {
			if item.GetName() == itemName {
				return position, index, item
			}
		}
	}

	return "", -1, nil
}

// AddLock adds lock to specific location
func (location *Location) AddLock(itemName string, locationToLock *Location) {
	location.locks = append(location.locks, NewLock(itemName, locationToLock))
}

// Unlock opens the door to another location
func (location *Location) Unlock() {
	location.locks = nil 	// key opens all doors in location
}

func (location *Location) addPlayerToLocation(player *Player) {
	location.mutex.Lock()
	defer location.mutex.Unlock()
	location.playersInLocation = append(location.playersInLocation, player)
}

func (location *Location) removePlayerFromLocation(player *Player) error {
	location.mutex.Lock()
	defer location.mutex.Unlock()
	
	index := -1
	for i, el := range location.playersInLocation {
		if (player.GetNickname() == el.GetNickname()) {
			index = i
			break
		}
	}

	if (index == -1) {
		return errors.New("No player with name " + player.GetNickname() + " in current location")
	}

	location.playersInLocation = append(location.playersInLocation[:index], location.playersInLocation[index + 1:]...)

	return nil
}

func (location *Location) getPlayersInLocation() []*Player {
	return location.playersInLocation
}
