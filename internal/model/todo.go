package model

type Todo struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	TagIds      []int  `json:"tag_ids"`
	IsCompleted bool   `json:"is_completed"`
}
