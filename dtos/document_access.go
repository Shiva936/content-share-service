package dtos

type DocumentsAccesses struct {
	DocumentsAccess []DocumentsAccess `json:"documents_access"`
}

type DocumentsAccess struct {
	ID         string `json:"id"`
	UserID     string `json:"user_id"`
	DocumentID string `json:"document_id"`
	AccessType string `json:"access_type"`
}
