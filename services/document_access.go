package services

import (
	"content-share/daos"
	"content-share/daos/models"
	"content-share/dtos"
	"context"
	"errors"
	"time"
)

func GetDocumentAccesses(ctx *context.Context, userId string, access *dtos.DocumentsAccess) (*dtos.DocumentsAccesses, error) {
	document, err := daos.DocumentsDB.GetDocumentById(access.DocumentID)
	if err != nil {
		return nil, err
	}
	if document.OwnerID != userId {
		return nil, errors.New("unauthorized: user is not the owner of document")
	}

	accesses := daos.DocumentAccessDB.GetDocumentAccessForDocument(access.DocumentID)
	var documentAccess []dtos.DocumentsAccess
	for i := range accesses {
		documentAccess = append(documentAccess, dtos.DocumentsAccess{
			ID:         accesses[i].ID,
			UserID:     accesses[i].UserID,
			DocumentID: accesses[i].DocumentID,
			AccessType: accesses[i].AccessType,
		})
	}
	return &dtos.DocumentsAccesses{
		DocumentsAccess: documentAccess,
	}, nil
}

func GrantDocumentAccess(ctx *context.Context, userId string, access *dtos.DocumentsAccess) (*dtos.DocumentsAccess, error) {
	document, err := daos.DocumentsDB.GetDocumentById(access.DocumentID)
	if err != nil {
		return nil, err
	}
	if document.OwnerID != userId {
		return nil, errors.New("unauthorized: user is not the owner of document")
	}

	daos.DocumentAccessDB.CreateDocumentAccess(models.DocumentsAccess{
		ID:         access.ID,
		UserID:     access.UserID,
		DocumentID: access.DocumentID,
		AccessType: access.AccessType,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	})

	return access, nil
}

func UpdateDocumentAccess(ctx *context.Context, userId string, access *dtos.DocumentsAccess) (*dtos.DocumentsAccess, error) {
	document, err := daos.DocumentsDB.GetDocumentById(access.DocumentID)
	if err != nil {
		return nil, err
	}
	if document.OwnerID != userId {
		return nil, errors.New("unauthorized: user is not the owner of document")
	}

	err = daos.DocumentAccessDB.UpdateDocumentAccess(models.DocumentsAccess{
		ID:         access.ID,
		UserID:     access.UserID,
		DocumentID: access.DocumentID,
		AccessType: access.AccessType,
		UpdatedAt:  time.Now(),
	})
	if err != nil {
		return nil, err
	}
	return access, nil
}

func DeleteDocumentAccess(ctx *context.Context, userId string, access *dtos.DocumentsAccess) error {
	document, err := daos.DocumentsDB.GetDocumentById(access.DocumentID)
	if err != nil {
		return err
	}
	if document.OwnerID != userId {
		return errors.New("unauthorized: user is not the owner of document")
	}
	return daos.DocumentAccessDB.DeleteDocumentAccess(models.DocumentsAccess{
		ID:         access.ID,
		DocumentID: access.DocumentID,
	})
}
