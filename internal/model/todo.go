package model

type ToDoModel struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Text string `json:"text"`
}