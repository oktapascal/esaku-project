package services

import (
	"context"
	"database/sql"
	"esaku-project/app/settings/models/domain"
	"esaku-project/app/settings/models/web"
	"esaku-project/app/settings/repository"
	"esaku-project/exceptions"
	"esaku-project/helpers"
	"github.com/go-playground/validator/v10"
)

type KelompokMenuServiceImpl struct {
	KelompokMenuRepository repository.KelompokMenuRepository
	Db                     *sql.DB
	Validate               *validator.Validate
}

func NewKelompokMenuServiceImpl(kelompokMenuRepository repository.KelompokMenuRepository, db *sql.DB, validate *validator.Validate) *KelompokMenuServiceImpl {
	return &KelompokMenuServiceImpl{
		KelompokMenuRepository: kelompokMenuRepository,
		Db:                     db,
		Validate:               validate,
	}
}

func (service *KelompokMenuServiceImpl) Save(ctx context.Context, request web.KelompokMenuSaveRequest) web.KelompokMenuResponse {
	err := service.Validate.Struct(request)

	helpers.PanicIfError(err)

	tx, err := service.Db.Begin()

	defer helpers.CommitOrRollback(tx, err)

	klpMenu := domain.KelompokMenu{
		KodeKlp: request.KodeKlp,
		Nama:    request.Nama,
	}

	klpMenu = service.KelompokMenuRepository.Save(ctx, tx, klpMenu)

	return web.ToKelompokMenuResponse(klpMenu)
}

func (service *KelompokMenuServiceImpl) Update(ctx context.Context, request web.KelompokMenuUpdateRequest) web.KelompokMenuResponse {
	err := service.Validate.Struct(request)

	helpers.PanicIfError(err)

	tx, err := service.Db.Begin()

	defer helpers.CommitOrRollback(tx, err)

	klpMenu, err := service.KelompokMenuRepository.FindById(ctx, tx, request.KodeKlp)

	if err != nil {
		panic(exceptions.NewErrorNotFound(err.Error()))
	}

	klpMenu.Nama = request.Nama

	klpMenu = service.KelompokMenuRepository.Update(ctx, tx, klpMenu)

	return web.ToKelompokMenuResponse(klpMenu)
}

func (service *KelompokMenuServiceImpl) Delete(ctx context.Context, kodeKlp string) {
	tx, err := service.Db.Begin()

	defer helpers.CommitOrRollback(tx, err)

	klpMenu, err := service.KelompokMenuRepository.FindById(ctx, tx, kodeKlp)

	if err != nil {
		panic(exceptions.NewErrorNotFound(err.Error()))
	}

	service.KelompokMenuRepository.Delete(ctx, tx, klpMenu)
}

func (service *KelompokMenuServiceImpl) FindById(ctx context.Context, kodeKlp string) web.KelompokMenuResponse {
	tx, err := service.Db.Begin()

	defer helpers.CommitOrRollback(tx, err)

	klpMenu, err := service.KelompokMenuRepository.FindById(ctx, tx, kodeKlp)

	if err != nil {
		panic(exceptions.NewErrorNotFound(err.Error()))
	}

	return web.ToKelompokMenuResponse(klpMenu)
}

func (service *KelompokMenuServiceImpl) FindAll(ctx context.Context) []web.KelompokMenuResponse {
	tx, err := service.Db.Begin()

	defer helpers.CommitOrRollback(tx, err)

	klpMenus := service.KelompokMenuRepository.FindAll(ctx, tx)

	return web.ToKelompokMenuResponses(klpMenus)
}
