package entity

type Video struct {
	Title       string `json:"title" binding:"min=2,max=100" validate="is-cool"`
	Description string `json:"description" binding:"min=2"`
	URL         string `json:"url" binding:"required,url`
	Author      Person `json:"author" binding:"required"`
}
