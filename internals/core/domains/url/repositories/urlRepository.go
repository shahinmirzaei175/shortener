package repositories

import (
	"log"
	"shortener/internals/adapters/framework/driven/db"
	"shortener/internals/config"
	"shortener/internals/core/domains/url/models"
)

type URLRepositoryInterface interface {
	GetAll() ([]models.Url, error)
	Store(book models.Url) (models.Url, error)
	FindById(id int) (models.Url, error)
	FindByShortLink(short string) (models.Url, error)
}

type urlRepository struct {
}

func URLRepository() URLRepositoryInterface {
	return urlRepository{}
}

func (ur urlRepository) GetAll() ([]models.Url, error) {
	var urls []models.Url

	res := db.DB.Find(&urls)
	if res.Error != nil {
		log.Fatal("err")
	}
	return urls, nil
}
func (ur urlRepository) Store(url models.Url) (models.Url, error) {

	res := db.DB.Create(&url)
	if res.Error != nil {
		log.Fatal("err")
	}
	return url, nil
}
func (ur urlRepository) FindById(id int) (models.Url, error) {
	var url models.Url
	result := db.DB.First(&url, "id = ?", id)
	if result.Error != nil {
		log.Fatal("err")
	}
	return url, nil
}
func (ur urlRepository) FindByShortLink(short string) (models.Url, error) {
	var url models.Url
	result := db.DB.First(&url, "short_url = ?", config.AppConfig.Domain+"/s/"+short)
	if result.Error != nil {
		log.Fatal("err")
	}
	return url, nil
}
