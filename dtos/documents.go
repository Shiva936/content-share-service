package dtos

type Documents struct {
	Documents []Document `json:"documents"`
}

type Document struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Content  string `json:"content"`
	OwnerID  string `json:"owner_id"`
	EditedBy string `json:"edited_by"`
}