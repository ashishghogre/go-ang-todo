package main

import (
	"fmt"
)

type todoItem struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Details string `json:"details"`
	CreatedTime string `json:"createdTime"`
	
	//Completed bool `json:"completed"`
}

func (item *todoItem) GoString() string{
	return fmt.Sprintf("ID: %d , Title: %s", item.ID, item.Title)
}

type createResponse struct {
	ID int `json:"id"`
}
