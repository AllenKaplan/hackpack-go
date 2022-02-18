package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Name        string `json:"company,omitempty"`
	Description string `json:"position,omitempty"`
	Status      string `json:"status,omitempty"`
}
