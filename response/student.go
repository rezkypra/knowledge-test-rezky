package response

import ()

type IndexStudent struct {
	ID string `json:"id"`
	StudentID string `json:"student_id"`
	Name string `json:"name"`
}

type ShowStudent struct {
	ID string `json:"id"`
	StudentID string `json:"student_id"`
	Name string `json:"name"`
	Gender string `json:"gender"`
	Address string `json:"address"`
	EntryYear uint `json:"entry_year"`
}