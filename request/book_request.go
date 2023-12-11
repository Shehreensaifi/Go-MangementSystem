package request

type BookRequest struct {
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
	Year        int    `json:"year"`
}
