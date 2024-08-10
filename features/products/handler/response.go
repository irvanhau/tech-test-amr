package handler

import (
	"github.com/google/uuid"
	"time"
)

type InputResponse struct {
	ID          uuid.UUID `json:"id" form:"id"`
	Name        string    `json:"name" form:"name"`
	Description string    `json:"description" form:"description"`
	CategoryID  string    `json:"category_id" form:"category_id"`
	CreatedAt   time.Time `json:"created_at" form:"created_at"`
}
