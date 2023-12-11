package response

type BookResponse struct {
	ID          int    `json:"ID"`
	Name        string `json:"name,omitempty"`
	Author      string `json:"author,omitempty"`
	Publication string `json:"publication,omitempty"`
	Year        int    `json:"year,omitempty"`
}
