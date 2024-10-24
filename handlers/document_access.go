package handlers

import (
	"content-share/dtos"
	"content-share/services"
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetDocumentAccesses(c *gin.Context) {
	userId := c.Param("userId")
	if userId == "" {
		err := errors.New("param : userId not provided")
		log.Println(err)
		c.JSON(http.StatusBadGateway, err)
		return
	}

	documentId := c.Param("documentId")
	if documentId == "" {
		err := errors.New("param : documentID not provided")
		log.Println(err)
		c.JSON(http.StatusBadGateway, err)
		return
	}

	requestContext := c.Request.Context()
	ctx, cancel := context.WithTimeout(requestContext, time.Duration(time.Second*30))
	defer cancel()
	c.Request = c.Request.WithContext(ctx)

	res, err := services.GetDocumentAccesses(&requestContext, userId, &dtos.DocumentsAccess{DocumentID: documentId})
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func GrantDocumentAccess(c *gin.Context) {
	userId := c.Param("userId")
	if userId == "" {
		err := errors.New("param : userId not provided")
		log.Println(err)
		c.JSON(http.StatusBadGateway, err)
		return
	}

	documentId := c.Param("documentId")
	if documentId == "" {
		err := errors.New("param : documentID not provided")
		log.Println(err)
		c.JSON(http.StatusBadGateway, err)
		return
	}

	requestContext := c.Request.Context()
	ctx, cancel := context.WithTimeout(requestContext, time.Duration(time.Second*30))
	defer cancel()
	c.Request = c.Request.WithContext(ctx)

	req := &dtos.DocumentsAccess{}
	if err := c.BindJSON(&req); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	req.DocumentID = documentId

	res, err := services.GrantDocumentAccess(&requestContext, userId, req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func UpdateDocumentAccess(c *gin.Context) {
	userId := c.Param("userId")
	if userId == "" {
		err := errors.New("param : userId not provided")
		log.Println(err)
		c.JSON(http.StatusBadGateway, err)
		return
	}

	documentId := c.Param("documentId")
	if documentId == "" {
		err := errors.New("param : documentID not provided")
		log.Println(err)
		c.JSON(http.StatusBadGateway, err)
		return
	}

	requestContext := c.Request.Context()
	ctx, cancel := context.WithTimeout(requestContext, time.Duration(time.Second*30))
	defer cancel()
	c.Request = c.Request.WithContext(ctx)

	req := &dtos.DocumentsAccess{}
	if err := c.BindJSON(&req); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if req.ID == "" {
		err := errors.New("request : DocumentAccessID not provided")
		log.Println(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	req.DocumentID = documentId

	res, err := services.UpdateDocumentAccess(&requestContext, userId, req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func DeleteDocumentAccess(c *gin.Context) {
	userId := c.Param("userId")
	if userId == "" {
		err := errors.New("param : UserID not provided")
		log.Println(err)
		c.JSON(http.StatusBadGateway, err)
		return
	}

	documentId := c.Param("documentId")
	if documentId == "" {
		err := errors.New("param : DocumentID not provided")
		log.Println(err)
		c.JSON(http.StatusBadGateway, err)
		return
	}

	documentAccessId := c.Param("documentAccessId")
	if documentAccessId == "" {
		err := errors.New("param : DocumentAccessID not provided")
		log.Println(err)
		c.JSON(http.StatusBadGateway, err)
		return
	}

	requestContext := c.Request.Context()
	ctx, cancel := context.WithTimeout(requestContext, time.Duration(time.Second*30))
	defer cancel()
	c.Request = c.Request.WithContext(ctx)

	err := services.DeleteDocumentAccess(&requestContext, userId, &dtos.DocumentsAccess{ID: documentAccessId, DocumentID: documentId})
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "SUCCESS",
		"message": "Resource Deleted Successfully",
	})
}
