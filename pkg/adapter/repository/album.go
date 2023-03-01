package repository

import (
	"github.com/go-gin-api/pkg/domain/model"
	"github.com/go-gin-api/pkg/usecase/repository"
	"github.com/jinzhu/gorm"
)

type albumRepository struct {
	db *gorm.DB
}

func NewAlbumRepository(db *gorm.DB) repository.AlbumRepository {
	return &albumRepository{db}
}

func (ar *albumRepository) FindAll(a []*model.Album) ([]*model.Album, error) {
	err := ar.db.Find(&a).Error

	if err != nil {
		return nil, err
	}

	return a, nil
}
