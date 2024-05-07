package books

import (
	"log"

	"gorm.io/gorm"
)

func GetBook(db *gorm.DB, id int) (*Book, error) {
	var b Book
	result := db.First(&b, id)
	if err := result.Error; err != nil {
		log.Fatalf("Error getting book: %v", err)
		return nil, err
	}

	return &b, nil
}

func GetBooks(db *gorm.DB) ([]Book, error) {
	var bs []Book
	result := db.Find(&bs)
	if err := result.Error; err != nil {
		log.Fatalf("Error getting books: %v", err)
		return nil, err
	}

	return bs, nil
}

func SearchBook(db *gorm.DB, name string) ([]Book, error) {
	var bs []Book
	result := db.Where("name = ?", name).Order("author DESC").Find(&bs)
	if err := result.Error; err != nil {
		log.Fatalf("Error searching book: %v", err)
		return nil, err
	}

	return bs, nil
}

func CreateBook(db *gorm.DB, book *Book) error {
	result := db.Create(book)
	if err := result.Error; err != nil {
		log.Fatalf("Error creating book: %v", err)
		return err
	}

	log.Println("Book created successfully")
	return nil
}

func UpdateBook(db *gorm.DB, book *Book) error {
	result := db.Save(book)
	if err := result.Error; err != nil {
		log.Fatalf("Error updating book: %v", err)
		return err
	}

	log.Println("Book updated successfully")
	return nil

}

func DeleteBook(db *gorm.DB, id int) error {
	result := db.Delete(&Book{}, id)
	if err := result.Error; err != nil {
		log.Fatalf("Error deleting book: %v", err)
		return err
	}

	log.Println("Book deleted successfully")
	return nil
}
