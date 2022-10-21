package services

import (
	"context"
	"database/sql"
	"esaku-project/app/settings/models/web"
	"esaku-project/app/settings/repository"
	"esaku-project/exceptions"
	"esaku-project/helpers"
	"github.com/go-playground/validator/v10"
)

type MenuServiceImpl struct {
	KelompokMenuRepository repository.KelompokMenuRepository
	MenuRepository         repository.MenuRepository
	Db                     *sql.DB
	Validate               *validator.Validate
}

func NewMenuServiceImpl(kelompokMenuRepository repository.KelompokMenuRepository, menuRepository repository.MenuRepository, db *sql.DB, validate *validator.Validate) *MenuServiceImpl {
	return &MenuServiceImpl{
		KelompokMenuRepository: kelompokMenuRepository,
		MenuRepository:         menuRepository,
		Db:                     db,
		Validate:               validate,
	}
}

func (service *MenuServiceImpl) Save(ctx context.Context, request web.MenuSaveRequest) {
	//TODO implement me
	panic("implement me")
}

func (service *MenuServiceImpl) Delete(ctx context.Context, klpMenu string) {
	//TODO implement me
	panic("implement me")
}

func (service *MenuServiceImpl) FindById(ctx context.Context, kodeKlp string) web.MenuResponse {
	tx, err := service.Db.Begin()
	defer helpers.CommitOrRollback(tx, err)

	klpMenu, err := service.KelompokMenuRepository.FindById(ctx, tx, kodeKlp)
	if err != nil {
		panic(exceptions.NewErrorNotFound(err.Error()))
	}

	menus := service.MenuRepository.FindById(ctx, tx, kodeKlp)

	return web.ToMenuResponse(klpMenu, menus)
}
