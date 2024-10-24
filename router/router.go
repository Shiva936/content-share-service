package router

import (
	"content-share/handlers"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	router := gin.Default()

	router.Use(RequestLogger())

	router.GET("v1/users", handlers.GetUsers)
	router.POST("v1/users", handlers.CreateUser)
	router.PUT("v1/users", handlers.UpdateUser)

	router.GET("v1/users/:userId/documents", handlers.GetUserDocuments)
	router.POST("v1/users/:userId/documents", handlers.CreateDocument)
	router.PUT("v1/users/:userId/documents", handlers.UpdateDocument)
	router.GET("v1/users/:userId/documents/:documentId", handlers.GetDocumentById)
	router.DELETE("v1/users/:userId/documents/:documentId", handlers.DeleteDocument)

	router.GET("v1/users/:userId/documents/:documentId/list-access", handlers.GetDocumentAccesses)
	router.POST("v1/users/:userId/documents/:documentId/grant-access", handlers.GrantDocumentAccess)
	router.PUT("v1/users/:userId/documents/:documentId/update-access", handlers.UpdateDocumentAccess)
	router.DELETE("v1/users/:userId/documents/:documentId/remove-access/:documentAccessId", handlers.DeleteDocumentAccess)

	return router
}
