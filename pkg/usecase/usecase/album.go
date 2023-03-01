package usecase

import (
	"github.com/go-gin-api/pkg/domain/model"
	"github.com/go-gin-api/pkg/usecase/repository"
)

type albumUsecase struct {
	albumRepository repository.AlbumRepository
	dBRepository    repository.DBRepository
}

type Album interface {
	List(u []*model.Album) ([]*model.Album, error)
}

func NewAlbumUsecase(r repository.AlbumRepository, d repository.DBRepository) Album {
	return &albumUsecase{r, d}
}

func (uu *albumUsecase) List(u []*model.Album) ([]*model.Album, error) {
	u, err := uu.albumRepository.FindAll(u)
	if err != nil {
		return nil, err
	}

	return u, nil
}
