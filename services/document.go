package services

import (
	"content-share/daos"
	"content-share/daos/models"
	"content-share/dtos"
	"context"
	"errors"
	"strings"
	"time"

	"golang.org/x/exp/maps"
)

func GetUserDocuments(ctx *context.Context, userId string) (*dtos.Documents, error) {
	documentMap := make(map[string]dtos.Document)
	documents := daos.DocumentsDB.GetDocumentForUser(userId)
	for i := range documents {
		documentMap[documents[i].ID] = dtos.Document{
			ID:       documents[i].ID,
			Name:     documents[i].Name,
			Content:  documents[i].Content,
			OwnerID:  documents[i].OwnerID,
			EditedBy: documents[i].EditedBy,
		}
	}
	documentAccesses := daos.DocumentAccessDB.GetDocumentAccessForUser(userId, "read")
	var accessDocumentIds []string
	for i := range documentAccesses {
		accessDocumentIds = append(accessDocumentIds, documentAccesses[i].DocumentID)
	}
	accessDocuments := daos.DocumentsDB.GetDocumentByIds(accessDocumentIds)
	for i := range accessDocuments {
		documentMap[accessDocuments[i].ID] = dtos.Document{
			ID:       accessDocuments[i].ID,
			Name:     accessDocuments[i].Name,
			Content:  accessDocuments[i].Content,
			OwnerID:  accessDocuments[i].OwnerID,
			EditedBy: accessDocuments[i].EditedBy,
		}
	}
	return &dtos.Documents{
		Documents: maps.Values(documentMap),
	}, nil
}

func GetUserDocumentById(ctx *context.Context, userId string, documentId string) (*dtos.Document, error) {
	document, err := daos.DocumentsDB.GetDocumentById(documentId)
	if err != nil {
		return nil, err
	}
	if document.OwnerID != userId {
		accesses := daos.DocumentAccessDB.GetDocumentAccessForDocument(document.ID)
		var hasWriteAccess bool
		for i := range accesses {
			if accesses[i].UserID == userId && strings.Contains(accesses[i].AccessType, "read") {
				hasWriteAccess = true
				break
			}
		}
		if !hasWriteAccess {
			return nil, errors.New("unauthorized : No read access")
		}
	}
	return &dtos.Document{
		ID:       document.ID,
		Name:     document.Name,
		Content:  document.Content,
		OwnerID:  document.OwnerID,
		EditedBy: document.EditedBy,
	}, nil
}

func CreateDocument(ctx *context.Context, document *dtos.Document) (*dtos.Document, error) {
	document.EditedBy = document.OwnerID
	daos.DocumentsDB.CreateDocument(models.Document{
		ID:        document.ID,
		Name:      document.Name,
		Content:   document.Content,
		OwnerID:   document.OwnerID,
		EditedBy:  document.EditedBy,
		CreatedAt: time.Now(),
	})
	return document, nil
}

func UpdateDocument(ctx *context.Context, userId string, document *dtos.Document) (*dtos.Document, error) {
	currentDocument, err := daos.DocumentsDB.GetDocumentById(document.ID)
	if err != nil {
		return nil, err
	}
	if currentDocument.OwnerID != userId {
		accesses := daos.DocumentAccessDB.GetDocumentAccessForDocument(document.ID)
		var hasWriteAccess bool
		for i := range accesses {
			if accesses[i].UserID == userId && strings.Contains(accesses[i].AccessType, "write") {
				hasWriteAccess = true
				break
			}
		}
		if !hasWriteAccess {
			return nil, errors.New("unauthorized : No write access")
		}
	}
	currentDocument.EditedBy = userId
	currentDocument.Content = document.Content
	currentDocument.Name = document.Name
	currentDocument.UpdatedAt = time.Now()

	err = daos.DocumentsDB.UpdateDocument(currentDocument)
	if err != nil {
		return nil, err
	}
	return document, nil
}

func DeleteDocument(ctx *context.Context, document *dtos.Document) error {
	return daos.DocumentsDB.DeleteDocument(models.Document{
		ID:      document.ID,
		OwnerID: document.OwnerID,
	})
}
