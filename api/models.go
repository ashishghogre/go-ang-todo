package main

type todoItem struct {
	Id int `json:"id"`
	Title string `json:"title"`
	//Completed bool `json:"completed"`
}

type createResponse struct {
	Id int `json:"id"`
}