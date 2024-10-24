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

func GetUsers(c *gin.Context) {

	requestContext := c.Request.Context()
	ctx, cancel := context.WithTimeout(requestContext, time.Duration(time.Second*30))
	defer cancel()
	c.Request = c.Request.WithContext(ctx)

	res, err := services.GetUsers(&requestContext)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func CreateUser(c *gin.Context) {

	requestContext := c.Request.Context()
	ctx, cancel := context.WithTimeout(requestContext, time.Duration(time.Second*30))
	defer cancel()
	c.Request = c.Request.WithContext(ctx)

	req := &dtos.User{}
	if err := c.BindJSON(&req); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	res, err := services.CreateUser(&requestContext, req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func UpdateUser(c *gin.Context) {

	requestContext := c.Request.Context()
	ctx, cancel := context.WithTimeout(requestContext, time.Duration(time.Second*30))
	defer cancel()
	c.Request = c.Request.WithContext(ctx)

	req := &dtos.User{}
	if err := c.BindJSON(&req); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if req.ID == "" {
		err := errors.New("request : Id not provided")
		log.Println(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	res, err := services.UpdateUser(&requestContext, req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, res)
}
