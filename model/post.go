package model

import "time"

type Post struct {
	UserId      string `json:"userId"`
	Title       string `json:"title" validate:"required,min=10,max=100"`
	Published   bool   `json:"published"`
	Description string `json:"description" validate:"required,min=10,max=1000"`
}

type PostUpdate struct {
	UserId      string `json:"userId"`
	Id          string `json:"id"`
	Title       string `json:"title" validate:"required,min=10,max=100"`
	Published   bool   `json:"published"`
	Description string `json:"description" validate:"required,min=10,max=1000"`
}

type PostResponse struct {
	Id          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Published   bool      `json:"published"`
	Created     time.Time `json:"created"`
}
