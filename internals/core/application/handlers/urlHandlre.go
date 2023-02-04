package handlers

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"shortener/internals/core/domains/url/models"
	"shortener/internals/core/domains/url/repositories"
	Utility "shortener/internals/utilities"
	"shortener/internals/utilities/response"
	"strconv"
	"time"
)

type UrlHandlerInterface interface {
	Index(c echo.Context) error
	Store(c echo.Context) error
	Show(c echo.Context) error
	Redirect(c echo.Context) error
}

type urlHandler struct {
	urlRepository repositories.URLRepositoryInterface
}

func UrlHandler() UrlHandlerInterface {
	return urlHandler{
		urlRepository: repositories.URLRepository(),
	}
}
func NewUrlHandler(repo repositories.URLRepositoryInterface) UrlHandlerInterface {
	return urlHandler{
		urlRepository: repo,
	}
}

func (uh urlHandler) Redirect(c echo.Context) error {
	short := c.Param("short")
	uh.urlRepository = repositories.URLRepository()
	show, err := uh.urlRepository.FindByShortLink(short)
	if err != nil {
		return err
	}
	return c.Redirect(301, show.MainUrl)
}

func (uh urlHandler) Index(c echo.Context) error {

	res, err := uh.urlRepository.GetAll()
	if err != nil {
		log.Fatal("err")
	}
	return c.JSON(http.StatusOK, response.Success(res, ""))
}
func (uh urlHandler) Store(c echo.Context) error {
	urlStoreRequest := new(models.UrlStoreRequest)
	if err := c.Bind(urlStoreRequest); err != nil {
		return c.JSON(http.StatusBadRequest, response.Fail(nil, "Data not found"))
	}

	if err := c.Validate(urlStoreRequest); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.Fail(nil, err.Error()))
	}
	book := models.Url{
		MainUrl:   Utility.ValidMainLink(urlStoreRequest.MainUrl),
		ShortUrl:  Utility.GenerateShortLink(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	res, _ := uh.urlRepository.Store(book)
	return c.JSON(http.StatusOK, response.Success(res, ""))
}
func (uh urlHandler) Show(c echo.Context) error {
	urlId := c.Param("id")
	atoi, err := strconv.Atoi(urlId)
	if err != nil {
		return err
	}
	show, err := uh.urlRepository.FindById(atoi)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, response.Success(show, ""))
}
