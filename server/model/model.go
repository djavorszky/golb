package model

import "time"

type BlogEntry struct {
	ID          string    `json:"id"`
	Author      string    `json:"author"`
	Tags        []string  `json:"tags"`
	Categories  []string  `json:"categories"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	CreateDate  time.Time `json:"createDate"`
	UpdateDate  time.Time `json:"updateDate"`
	PublishDate time.Time `json:"publishDate"`
}
