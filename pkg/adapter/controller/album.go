package controller

import (
	"net/http"

	"github.com/go-gin-api/pkg/domain/model"
	"github.com/go-gin-api/pkg/usecase/usecase"
)

type albumController struct {
	albumUsecase usecase.Album
}

type Album interface {
	GetAlbums(ctx Context) error
}

func NewAlbumController(us usecase.Album) Album {
	return &albumController{us}
}

func (uc *albumController) GetAlbums(ctx Context) error {
	var u []*model.Album

	u, err := uc.albumUsecase.List(u)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, u)
}
