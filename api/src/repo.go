package main

import (
	"fmt"
	"time"
	"strconv"
)

func incrementID(ids chan int) {
	currentID, _ := strconv.Atoi(dbClient.getMaxId())
	for {
		currentID++
		ids <- currentID
		dbClient.updateMaxId(strconv.Itoa(currentID))
	}
}

func createItemInDb(item todoItem) int {
	id := <-ids
	fmt.Printf("getting id %d", id)
	item.ID = id
	item.CreatedTime = time.Now().Format("02-01-2006 15:04:05")
	dbClient.upsertItem(strconv.Itoa(id), item)
	return id
}

func updateItemInDb(id string, item todoItem) {
	dbClient.upsertItem(id, item)
}

func getItemsFromDb() []todoItem {
	var items []todoItem
	items = dbClient.getItems()
	if items == nil {
		items = []todoItem{}
	}
	return items
}

func getItemFromDb(id string) todoItem {
	var item todoItem
	item = dbClient.getItem(id)
	return item
}
