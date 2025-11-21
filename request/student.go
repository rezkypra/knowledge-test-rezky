package request

import ()

type StoreStudent struct {
	StudentID string `json:"student_id" binding:"required"`
	Name      string `json:"name" binding:"required"`
	Gender    string `json:"gender" binding:"required"`
	Address   string `json:"address" binding:"required"`
	EntryYear uint   `json:"entry_year" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
}

type UpdateStudent struct {
	Name      string `json:"name"`
	Gender    string `json:"gender"`
	Address   string `json:"address"`
	EntryYear uint   `json:"entry_year"`
	Email     string `json:"email" binding:"omitempty,email"`
}