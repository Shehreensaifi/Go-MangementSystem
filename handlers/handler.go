package handlers

import (
	"book-store/request"
	"book-store/services"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

type BookHandler struct {
	BookService services.SvcInterface
}

func New(service services.SvcInterface) BookHandler {
	return BookHandler{BookService: service}
}

func (b BookHandler) Create(c *gin.Context) {

	var req request.BookRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Println("Error while binding json.", err)
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	resp, err := b.BookService.Create(c, req)
	if err != nil {
		panic(err)
	}
	panic(errors.New("shit"))

	c.JSON(http.StatusCreated, resp)
}

func (b BookHandler) FindById(c *gin.Context) {
	id := c.Param("id")
	idVal, err := strconv.Atoi(id)
	if err != nil {
		log.Println("Invalid ID in params.")
		panic(err)
	}

	bookResp, err := b.BookService.FindById(c, idVal)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, nil)
			return
		}
		panic(err)
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, bookResp)
}

func (b BookHandler) FindAll(c *gin.Context) {
	resp, err := b.BookService.FindAll(c)
	if err != nil {
		panic(err)
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, resp)

}

func (b BookHandler) Update(c *gin.Context) {
	bookReq := request.BookRequest{}
	err := c.ShouldBindJSON(&bookReq)
	if err != nil {
		panic(err)
	}

	idInt := c.Param("id")
	id, err := strconv.Atoi(idInt)
	if err != nil {
		log.Println("Invalid ID in params.")
		panic(err)
	}

	resp, err := b.BookService.Update(c, id, bookReq)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, nil)
			return
		}
		panic(err)
	}

	c.JSON(http.StatusOK, resp)
}

func (b BookHandler) Delete(c *gin.Context) {
	idInt := c.Param("id")
	id, err := strconv.Atoi(idInt)
	if err != nil {
		log.Println("Invalid ID in params.")
		panic(err)
	}

	err = b.BookService.Delete(c, id)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusNoContent, nil)

}
