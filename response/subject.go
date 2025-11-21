package response

type Subject struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Credits uint8  `json:"credits"`
}
