package models

import "time"

// Models to replicate DB table structure

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

type Document struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Content   string    `json:"content"`
	OwnerID   string    `json:"owner_id"`
	EditedBy  string    `json:"edited_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DocumentsAccess struct {
	ID         string    `json:"id"`
	UserID     string    `json:"user_id"`
	DocumentID string    `json:"document_id"`
	AccessType string    `json:"access_type"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
