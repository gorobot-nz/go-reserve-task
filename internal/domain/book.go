package domain

type Book struct {
	ID      int64    `json:"id" binding:"required"`
	Title   string   `json:"title" binding:"required"`
	Authors []string `json:"authors" binding:"required"`
	Year    string   `json:"year" binding:"required"`
}
