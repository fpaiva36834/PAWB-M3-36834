package entity

type Book struct {
	ID        uint64 `json:"id"`
	Title     string `json:"title"`
	Year      string `json:"year"`
	Stock     string `json:"stock"`
	BookCover string `json:"book_cover"`
	UserId    uint64 `json:"user_id"`
	User      User   `json:"user"`
}
