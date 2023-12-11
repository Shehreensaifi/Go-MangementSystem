package services

import (
	"book-store/datastores"
	"book-store/model"
	"book-store/request"
	"book-store/response"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"log"
)

type BookService struct {
	BookDatastore datastores.BookDataStoreInterface
}

func New(store datastores.BookDataStoreInterface) BookService {

	return BookService{BookDatastore: store}
}

func (b BookService) FindById(c *gin.Context, id int) (response.BookResponse, error) {

	bookModel, err := b.BookDatastore.FindById(c, id)
	if err != nil {
		return response.BookResponse{}, err
	}

	resp := transformModelToResponse(bookModel)
	return resp, nil
}

func (b BookService) FindAll(c *gin.Context) ([]response.BookResponse, error) {
	result, err := b.BookDatastore.FindAll(c)
	if err != nil {
		return nil, err
	}

	var resp []response.BookResponse
	for _, value := range result {
		bookData := transformModelToResponse(&value)
		resp = append(resp, bookData)
	}
	return resp, nil
}

func (b BookService) Update(c *gin.Context, id int, req request.BookRequest) (response.BookResponse, error) {
	bookModel, err := b.BookDatastore.FindById(c, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("Record with id %d not found", id)
			return response.BookResponse{}, err
		}
		return response.BookResponse{}, err
	}

	getUpdatedModel(c, id, &req, bookModel)
	respModel, err := b.BookDatastore.Update(c, bookModel)
	if err != nil {
		return response.BookResponse{}, err
	}

	resp := transformModelToResponse(respModel)
	return resp, nil
}

func (b BookService) Delete(c *gin.Context, id int) error {
	err := b.BookDatastore.Delete(c, id)
	if err != nil {
		return err
	}

	return nil
}

func (b BookService) Create(c *gin.Context, request request.BookRequest) (response.BookResponse, error) {

	bookModel := transformRequestToModel(request)

	bookModelResp, err := b.BookDatastore.Create(c, &bookModel)
	if err != nil {
		return response.BookResponse{}, err
	}

	bookResp := transformModelToResponse(bookModelResp)
	return bookResp, nil
}

func transformModelToResponse(model *model.Book) response.BookResponse {
	resp := response.BookResponse{
		ID:          int(model.ID),
		Name:        model.Name,
		Author:      model.Author,
		Publication: model.Publication,
		Year:        model.Year,
	}

	return resp
}

func transformRequestToModel(bookRequest request.BookRequest) model.Book {

	bookModel := model.Book{
		Name:        bookRequest.Name,
		Author:      bookRequest.Author,
		Publication: bookRequest.Publication,
		Year:        bookRequest.Year,
	}

	return bookModel
}

func getUpdatedModel(c *gin.Context, id int, from *request.BookRequest, to *model.Book) {
	if from.Author != "" {
		to.Author = from.Author
	}

	if from.Publication != "" {
		to.Publication = from.Publication
	}

	if from.Name != "" {
		to.Name = from.Name
	}

	if from.Year != 0 {
		to.Year = from.Year
	}

	to.ID = uint(id)
}
