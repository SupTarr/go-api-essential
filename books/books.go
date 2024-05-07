package books

import "github.com/SupTarr/go-api-essential/utils"

type Book struct {
	utils.BaseModel
	Name        string `json:"name"`
	Author      string `json:"author"`
	Description string `json:"description"`
}
