package handler

import (
	"errors"
	"net/http"

	"github.com/aliazizii/url-shotener/internal/model"
	"github.com/aliazizii/url-shotener/internal/requset"

	"github.com/aliazizii/url-shotener/internal/store"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Url struct {
	Store store.Url
}

type response struct {
	Link string `json:"link"`
}

func (iUrl Url) FindByHash(c echo.Context) error {
	var req requset.FindByHashReq
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "")
	}

	link, err := iUrl.Store.FindByHash(req.Hash)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, "")
		}
		return c.JSON(http.StatusInternalServerError, "")
	}

	response := response{
		Link: link,
	}

	return c.JSON(http.StatusOK, response)

}

func (iUrl Url) Insert(c echo.Context) error {
	var req requset.InsertReq
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "")
	}

	m := model.Url{
		Link: req.Link,
	}

	hash, err := iUrl.Store.Insert(m)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "")
	}

	response := response{
		Link: hash,
	}

	return c.JSON(http.StatusOK, response)

}
