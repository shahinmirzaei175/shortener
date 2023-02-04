package tests

import (
	"shortener/internals/core/domains/url/models"
	"shortener/internals/core/domains/url/repositories"
	"time"
)

type URLRepository struct{}

func (U URLRepository) GetAll() ([]models.Url, error) {
	data := []models.Url{
		{
			ID:        1,
			MainUrl:   "http://mainurl1.com",
			ShortUrl:  "http://s.com",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        1,
			MainUrl:   "http://mainurl1.com",
			ShortUrl:  "http://s.com",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        1,
			MainUrl:   "http://mainurl1.com",
			ShortUrl:  "http://s.com",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	return data, nil
}

func (U URLRepository) Store(book models.Url) (models.Url, error) {
	//TODO implement me
	panic("implement me")
}

func (U URLRepository) FindById(id int) (models.Url, error) {
	//TODO implement me
	panic("implement me")
}

func (U URLRepository) FindByShortLink(short string) (models.Url, error) {
	//TODO implement me
	panic("implement me")
}

func URLRepositoryMock() repositories.URLRepositoryInterface {
	return URLRepository{}
}
