package daos

import (
	"content-share/daos/models"
	"errors"
	"strings"
	"sync"

	"github.com/google/uuid"
)

type IDocuments interface {
	GetDocumentForUser(userId string) []models.Document
	GetDocumentById(documentId string) (models.Document, error)
	GetDocumentByIds(documentId []string) []models.Document
	CreateDocument(Document models.Document)
	UpdateDocument(Document models.Document) error
	DeleteDocument(Document models.Document) error
}

type Documents struct {
	sync.Mutex
	Documents []models.Document
}

func NewDocuments() IDocuments {
	return &Documents{
		Documents: make([]models.Document, 0),
	}
}

func (u *Documents) GetDocumentForUser(userId string) []models.Document {
	var documents []models.Document
	for i := range u.Documents {
		if u.Documents[i].OwnerID == userId {
			documents = append(documents, u.Documents[i])
		}
	}
	return documents
}

func (u *Documents) GetDocumentByIds(documentIds []string) []models.Document {
	if len(documentIds) == 0 {
		return nil
	}
	allDocumentIds := strings.Join(documentIds, ",")
	var documents []models.Document
	for i := range u.Documents {
		if strings.Contains(allDocumentIds, u.Documents[i].ID) {
			documents = append(documents, u.Documents[i])
		}
	}
	return documents
}

func (u *Documents) GetDocumentById(documentId string) (models.Document, error) {
	var document models.Document
	var hasFound bool
	for i := range u.Documents {
		if u.Documents[i].ID == documentId {
			hasFound = true
			document = u.Documents[i]
		}
	}
	if !hasFound {
		return document, errors.New("document not found")
	}
	return document, nil
}

func (u *Documents) CreateDocument(Document models.Document) {
	u.Lock()
	defer u.Unlock()

	Document.ID = uuid.NewString()
	u.Documents = append(u.Documents, Document)
}

func (u *Documents) UpdateDocument(Document models.Document) error {
	u.Lock()
	defer u.Unlock()

	var hasFound bool
	for i := range u.Documents {
		if u.Documents[i].ID == Document.ID {
			hasFound = true
			u.Documents[i] = Document
		}
	}
	if !hasFound {
		return errors.New("document not found")
	}
	return nil
}

func (u *Documents) DeleteDocument(Document models.Document) error {
	u.Lock()
	defer u.Unlock()

	indexFound := -1
	for i := range u.Documents {
		if u.Documents[i].ID == Document.ID {
			if u.Documents[i].OwnerID != Document.OwnerID {
				return errors.New("user not an owner of the document")
			}
			indexFound = i
			break
		}
	}
	if indexFound == -1 {
		return errors.New("document not found")
	}

	u.Documents = append(u.Documents[:indexFound], u.Documents[indexFound+1:]...)
	return nil
}
