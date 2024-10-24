package daos

import (
	"content-share/daos/models"
	"errors"
	"strings"
	"sync"

	"github.com/google/uuid"
)

type IDocumentAccesses interface {
	GetDocumentAccessForUser(userId string, accessType string) []models.DocumentsAccess
	GetDocumentAccessForDocument(documentId string) []models.DocumentsAccess
	CreateDocumentAccess(documentAccess models.DocumentsAccess)
	UpdateDocumentAccess(documentAccess models.DocumentsAccess) error
	DeleteDocumentAccess(documentAccess models.DocumentsAccess) error
}

type DocumentAccesses struct {
	sync.Mutex
	DocumentAccesses []models.DocumentsAccess
}

func NewDocumentAccess() IDocumentAccesses {
	return &DocumentAccesses{
		DocumentAccesses: make([]models.DocumentsAccess, 0),
	}
}

func (u *DocumentAccesses) GetDocumentAccessForUser(userId string, accessType string) []models.DocumentsAccess {
	var documentAccess []models.DocumentsAccess
	for i := range u.DocumentAccesses {
		if u.DocumentAccesses[i].UserID == userId && strings.Contains(u.DocumentAccesses[i].AccessType, accessType) {
			documentAccess = append(documentAccess, u.DocumentAccesses[i])
		}
	}
	return documentAccess
}

func (u *DocumentAccesses) GetDocumentAccessForDocument(documentId string) []models.DocumentsAccess {
	var documentAccess []models.DocumentsAccess
	for i := range u.DocumentAccesses {
		if u.DocumentAccesses[i].DocumentID == documentId {
			documentAccess = append(documentAccess, u.DocumentAccesses[i])
		}
	}
	return documentAccess
}

func (u *DocumentAccesses) CreateDocumentAccess(documentAccess models.DocumentsAccess) {
	u.Lock()
	defer u.Unlock()

	documentAccess.ID = uuid.NewString()
	u.DocumentAccesses = append(u.DocumentAccesses, documentAccess)
}

func (u *DocumentAccesses) UpdateDocumentAccess(documentAccess models.DocumentsAccess) error {
	u.Lock()
	defer u.Unlock()

	var hasFound bool
	for i := range u.DocumentAccesses {
		if u.DocumentAccesses[i].ID == documentAccess.ID && u.DocumentAccesses[i].DocumentID == documentAccess.DocumentID {
			hasFound = true
			u.DocumentAccesses[i] = documentAccess
		}
	}
	if !hasFound {
		return errors.New("DocumentAccess not found")
	}
	return nil
}

func (u *DocumentAccesses) DeleteDocumentAccess(documentAccess models.DocumentsAccess) error {
	u.Lock()
	defer u.Unlock()

	indexFound := -1
	for i := range u.DocumentAccesses {
		if u.DocumentAccesses[i].ID == documentAccess.ID && u.DocumentAccesses[i].DocumentID == documentAccess.DocumentID {
			indexFound = i
			break

		}
	}
	if indexFound == -1 {
		return errors.New("DocumentAccess not found")
	}

	u.DocumentAccesses = append(u.DocumentAccesses[:indexFound], u.DocumentAccesses[indexFound+1:]...)
	return nil
}
