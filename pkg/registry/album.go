package registry

import (
	"github.com/go-gin-api/pkg/adapter/controller"
	"github.com/go-gin-api/pkg/adapter/repository"
	"github.com/go-gin-api/pkg/usecase/usecase"
)

func (r *registry) NewAlbumController() controller.Album {
	u := usecase.NewAlbumUsecase(
		repository.NewAlbumRepository(r.db),
		repository.NewDBRepository(r.db),
	)

	return controller.NewAlbumController(u)
}
