package services

import (
	"book-store/request"
	"book-store/response"
	"github.com/gin-gonic/gin"
)

type SvcInterface interface {
	Create(c *gin.Context, request request.BookRequest) (response.BookResponse, error)
	FindById(c *gin.Context, id int) (response.BookResponse, error)
	FindAll(c *gin.Context) ([]response.BookResponse, error)
	Update(c *gin.Context, id int, req request.BookRequest) (response.BookResponse, error)
	Delete(c *gin.Context, id int) error
}
