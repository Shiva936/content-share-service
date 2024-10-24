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

func GetUserDocuments(c *gin.Context) {
	userId := c.Param("userId")
	if userId == "" {
		err := errors.New("param : UserID not provided")
		log.Println(err)
		c.JSON(http.StatusBadGateway, err)
		return
	}

	requestContext := c.Request.Context()
	ctx, cancel := context.WithTimeout(requestContext, time.Duration(time.Second*30))
	defer cancel()
	c.Request = c.Request.WithContext(ctx)

	res, err := services.GetUserDocuments(&requestContext, userId)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func GetDocumentById(c *gin.Context) {
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

	requestContext := c.Request.Context()
	ctx, cancel := context.WithTimeout(requestContext, time.Duration(time.Second*30))
	defer cancel()
	c.Request = c.Request.WithContext(ctx)

	res, err := services.GetUserDocumentById(&requestContext, userId, documentId)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func CreateDocument(c *gin.Context) {
	userId := c.Param("userId")
	if userId == "" {
		err := errors.New("param : UserID not provided")
		log.Println(err)
		c.JSON(http.StatusBadGateway, err)
		return
	}

	requestContext := c.Request.Context()
	ctx, cancel := context.WithTimeout(requestContext, time.Duration(time.Second*30))
	defer cancel()
	c.Request = c.Request.WithContext(ctx)

	req := &dtos.Document{}
	if err := c.BindJSON(&req); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	req.OwnerID = userId
	res, err := services.CreateDocument(&requestContext, req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func UpdateDocument(c *gin.Context) {
	userId := c.Param("userId")
	if userId == "" {
		err := errors.New("param : UserID not provided")
		log.Println(err)
		c.JSON(http.StatusBadGateway, err)
		return
	}

	requestContext := c.Request.Context()
	ctx, cancel := context.WithTimeout(requestContext, time.Duration(time.Second*30))
	defer cancel()
	c.Request = c.Request.WithContext(ctx)

	req := &dtos.Document{}
	if err := c.BindJSON(&req); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if req.ID == "" {
		err := errors.New("request : DocumentID not provided")
		log.Println(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	res, err := services.UpdateDocument(&requestContext, userId, req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func DeleteDocument(c *gin.Context) {
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

	requestContext := c.Request.Context()
	ctx, cancel := context.WithTimeout(requestContext, time.Duration(time.Second*30))
	defer cancel()
	c.Request = c.Request.WithContext(ctx)

	err := services.DeleteDocument(&requestContext, &dtos.Document{ID: documentId, OwnerID: userId})
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
