package dto

type BookRequest struct {
	Title         string `json:"title" form:"title" binding:"required"`
	Author        string `json:"author" form:"author" binding:"required"`
	ISBN          string `json:"isbn" form:"isbn" binding:"required"`
	PublishedDate string `json:"published_date" form:"published_date" binding:"required"`
}

type UpdateBookRequest struct {
	Title         string `json:"title" form:"title"`
	Author        string `json:"author" form:"author"`
	ISBN          string `json:"isbn" form:"isbn"`
	PublishedDate string `json:"published_date" form:"published_date"`
}
