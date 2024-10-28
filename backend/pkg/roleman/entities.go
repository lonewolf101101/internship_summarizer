package roleman

import "undrakh.net/summarizer/pkg/entities"

type Role struct {
	entities.Model
	RID         uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
