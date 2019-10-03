package main

import (
	"fmt"
	"strconv"
)

func incrementId(ids chan int){
	currentId, _ := strconv.Atoi(dbClient.getMaxId())
	for{
		fmt.Printf("Inserting id %d",currentId)
		currentId++
		ids <- currentId
		dbClient.updateMaxId(strconv.Itoa(currentId))
	}
}

func createItemInDb(item todoItem) int{
	id := <- ids
	fmt.Printf("getting id %d", id)
	item.Id = id
	dbClient.upsertItem(strconv.Itoa(id), item)
	return id
}

func updateItemInDb(id string, item todoItem){
	dbClient.upsertItem(id, item)
}

func getItemsFromDb() []todoItem{
	var items []todoItem
	items = dbClient.getItems()
	if items == nil{
		items = []todoItem{}
	}
	return items
}

func getItemFromDb(id string) todoItem{
	var item todoItem
	item = dbClient.getItem(id)
	return item
}