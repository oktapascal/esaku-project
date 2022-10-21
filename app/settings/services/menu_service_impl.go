package services

import (
	"context"
	"database/sql"
	"esaku-project/app/settings/models/web"
	"esaku-project/app/settings/repository"
	"github.com/go-playground/validator/v10"
)

type MenuServiceImpl struct {
	MenuRepository repository.MenuRepository
	Db             *sql.DB
	Validate       *validator.Validate
}

func NewMenuServiceImpl(menuRepository repository.MenuRepository, db *sql.DB, validate *validator.Validate) *MenuServiceImpl {
	return &MenuServiceImpl{
		MenuRepository: menuRepository,
		Db:             db,
		Validate:       validate,
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

func (service *MenuServiceImpl) FindById(ctx context.Context, klpMenu string) {
	//TODO implement me
	panic("implement me")
}
