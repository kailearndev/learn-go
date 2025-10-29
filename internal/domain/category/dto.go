package category

type CategoryRequest struct {
	Name        string   `json:"name" validate:"required,min=3,max=100"`
	ImageURLs   []string `json:"image_urls" validate:"omitempty,dive,url"`
	Description string   `json:"description,omitempty"`
	CategoryID  *string  `json:"category_id,omitempty"` // optional
}
