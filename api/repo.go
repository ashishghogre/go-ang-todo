package main

import (
	"fmt"
	"strconv"
)

func incrementId(ids chan int){
	currentId := 0
	for{
		fmt.Printf("Inserting id %d",currentId)
		currentId++
		ids <- currentId
	}
}

func createItemInDb(item todoItem){
	id := <- ids
	fmt.Printf("getting id %d", id)
	item.Id = id
	dbClient.upsertItem(strconv.Itoa(id), item)
}

func getItemsFromDb() []todoItem{
	var items []todoItem
	items = dbClient.getItems()
	if items == nil{
		items = []todoItem{}
	}
	return items
}