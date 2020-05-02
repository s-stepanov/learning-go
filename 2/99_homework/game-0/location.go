package main

import (
	"fmt"
)

type Location struct {
	name string
	description string
	nearbyLocations []*Location
	availableItems map[string][]*Item
}

func NewLocation(name string) *Location {
	location := new(Location)
	location.name = name
	location.description = ""
	location.availableItems = make(map[string][]*Item)
	
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

func (location *Location) AddNearbyLocation(locationToAdd *Location) error {
	if location.HasNearbyLocation(locationToAdd.name) || locationToAdd.name == location.name {
		return fmt.Errorf("Cannot add nearby location with name: " + location.GetName())
	}
	location.nearbyLocations = append(location.nearbyLocations, locationToAdd)
	return nil
}

func (location *Location) HasNearbyLocation(locationName string) bool {
	for _, location := range location.nearbyLocations {
		if location.name == locationName {
			return true
		}
	}
	return false
}

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

func (location *Location) AddItem(position string, itemName string, isWearable bool) {
	location.availableItems[position] = append(location.availableItems[position], NewItem(itemName, isWearable))
}

func (location *Location) GetAvailableItems() map[string][]*Item {
	return location.availableItems
}

func (location *Location) GetAvailableItemsString() (result string) {
	mapIndex := 0
	for itemPosition, items := range location.availableItems {
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

func (location *Location) FindItem(itemName string) (string, int, *Item) {
	fmt.Println(location)
	for position, items := range location.GetAvailableItems() {
		for index, item := range items {
			if item.GetName() == itemName {
				return position, index, item
			}
		}
	}

	return "", -1, nil
}
