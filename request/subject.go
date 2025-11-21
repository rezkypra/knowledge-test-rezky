package request

type CreateSubject struct {
	Name    string `json:"name" binding:"required"`
	Credits uint8  `json:"credits" binding:"required"`
}

type UpdateSubject struct {
	Name    string `json:"name"`
	Credits uint8  `json:"credits"`
}

type AttachSubject struct {
	StudentID string `json:"student_id" binding:"required"`
	SubjectID string `json:"subject_id" binding:"required"`
}
