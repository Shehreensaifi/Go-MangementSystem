package datastores

import (
	"book-store/model"
	"github.com/gin-gonic/gin"
)

type BookDataStoreInterface interface {
	Create(c *gin.Context, book *model.Book) (*model.Book, error)
	FindById(c *gin.Context, id int) (*model.Book, error)
	Update(c *gin.Context, book *model.Book) (*model.Book, error)
	FindAll(c *gin.Context) ([]model.Book, error)
	Delete(c *gin.Context, id int) error
}
