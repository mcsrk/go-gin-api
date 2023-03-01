package repository

import (
	"github.com/go-gin-api/pkg/domain/model"
)

type AlbumRepository interface {
	FindAll(u []*model.Album) ([]*model.Album, error)
}
