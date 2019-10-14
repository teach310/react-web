package model

type Todo struct {
	ID     int    `json:"id"`
	IsDone bool   `json:"isDone"`
	Name   string `json:"name"`
}
