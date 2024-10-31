package roleman

import "undrakh.net/summarizer/pkg/entities"

const (
	ROLE_BASIC uint = 2
	ROLE_ADMIN uint = 1
)

type Role struct {
	entities.Model        // Assuming entities.Model includes common fields
	RID            uint   `gorm:"primaryKey;autoIncrement;column:rid" json:"rid"` // Map to 'rid' column
	Name           string `json:"name"`
	Description    string `json:"description,omitempty"`
}
