package main

type Item struct {
	name string
	wearable bool
}

func NewItem(name string, isWearable bool) *Item {
	item := new(Item)
	item.name = name
	item.wearable = isWearable

	return item
}

func (item *Item) GetName() string {
	return item.name
}

func (item *Item) IsWearable() bool {
	return item.wearable
}

type Lock struct {
	Item
	locked bool
	locationToLock *Location
}

func NewLock(lockName string, locationToLock *Location) *Lock {
	lock := new(Lock)
	lock.name = lockName
	lock.locked = true
	lock.locationToLock = locationToLock

	return lock
}