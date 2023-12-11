package datastores

import (
	"book-store/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
)

type BookDataStore struct {
	db *gorm.DB
}

func New(db *gorm.DB) BookDataStore {
	db.AutoMigrate(&model.Book{})
	return BookDataStore{db: db}
}

func (b BookDataStore) FindById(c *gin.Context, id int) (*model.Book, error) {
	var book model.Book
	err := b.db.First(&book, id).Error
	if err != nil {
		log.Println("Error from Datastore: ", err)
		return nil, err
	}

	return &book, nil
}

func (b BookDataStore) Update(c *gin.Context, book *model.Book) (*model.Book, error) {

	err := b.db.Save(book).Error
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (b BookDataStore) FindAll(c *gin.Context) ([]model.Book, error) {
	var books []model.Book
	err := b.db.Find(&books).Error
	if err != nil {
		log.Println("Error from Datastore: ", err)
		return nil, err
	}
	return books, nil
}

func (b BookDataStore) Delete(c *gin.Context, id int) error {
	var tags model.Book
	result := b.db.Where("id = ?", id).Delete(&tags)

	if result.Error != nil {
		log.Println("Error from Datastore: ", result.Error)
		return result.Error
	}

	return nil
}

func (b BookDataStore) Create(c *gin.Context, book *model.Book) (*model.Book, error) {
	err := b.db.Create(book).Error
	if err != nil {
		log.Println("Error from Datastore: ", err)
		return nil, err
	}
	return book, nil
}
