package domain

type Book struct {
	ID      int64    `json:"id"`
	Title   string   `json:"title"`
	Authors []string `json:"authors"`
	Year    string   `json:"year"`
}
