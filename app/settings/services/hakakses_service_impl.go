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

type HakAksesServiceImpl struct {
	HakAksesRepository repository.HakAksesRepository
	Db                 *sql.DB
	Validate           *validator.Validate
}

func NewHakAksesServiceImpl(hakAksesRepository repository.HakAksesRepository, db *sql.DB, validate *validator.Validate) *HakAksesServiceImpl {
	return &HakAksesServiceImpl{
		HakAksesRepository: hakAksesRepository,
		Db:                 db,
		Validate:           validate,
	}
}

func (service *HakAksesServiceImpl) Save(ctx context.Context, request web.HakAksesSaveRequest) web.HakAksesListResponse {
	err := service.Validate.Struct(request)
	helpers.PanicIfError(err)

	tx, err := service.Db.Begin()
	defer helpers.CommitOrRollback(tx, err)

	password, err := helpers.HashPassword(request.Password)
	helpers.PanicIfError(err)

	karyawan := domain.Karyawan{
		Nik: request.Nik,
	}

	kelompokMenu := domain.KelompokMenu{
		KodeKlp: request.KelompokMenu,
	}

	akses := domain.HakAkses{
		Password:       password,
		KelompokAkses:  request.KelompokAkses,
		StatusAdmin:    request.StatusAdmin,
		DefaultProgram: request.DefaultProgram,
		Karyawan:       karyawan,
		KelompokMenu:   kelompokMenu,
	}

	akses = service.HakAksesRepository.Save(ctx, tx, akses)

	return web.ToHakAksesListResponse(akses)
}

func (service *HakAksesServiceImpl) Update(ctx context.Context, request web.HakAksesUpdateRequest) web.HakAksesListResponse {
	err := service.Validate.Struct(request)
	helpers.PanicIfError(err)

	tx, err := service.Db.Begin()
	defer helpers.CommitOrRollback(tx, err)

	akses, err := service.HakAksesRepository.FindById(ctx, tx, request.Nik)

	if err != nil {
		panic(exceptions.NewErrorNotFound(err.Error()))
	}

	akses.KodeKlp = request.KelompokMenu
	akses.StatusAdmin = request.StatusAdmin
	akses.KelompokAkses = request.KelompokAkses
	akses.DefaultProgram = request.DefaultProgram

	akses = service.HakAksesRepository.Update(ctx, tx, akses)

	return web.ToHakAksesListResponse(akses)
}

func (service *HakAksesServiceImpl) Delete(ctx context.Context, nik string) {
	tx, err := service.Db.Begin()
	defer helpers.CommitOrRollback(tx, err)

	akses, err := service.HakAksesRepository.FindById(ctx, tx, nik)

	if err != nil {
		panic(exceptions.NewErrorNotFound(err.Error()))
	}

	service.HakAksesRepository.Delete(ctx, tx, akses)
}

func (service *HakAksesServiceImpl) FindById(ctx context.Context, nik string) web.HakAksesDetailResponse {
	tx, err := service.Db.Begin()
	defer helpers.CommitOrRollback(tx, err)

	akses, err := service.HakAksesRepository.FindById(ctx, tx, nik)

	if err != nil {
		panic(exceptions.NewErrorNotFound(err.Error()))
	}

	return web.ToHakAksesDetailResponse(akses)
}

func (service *HakAksesServiceImpl) FindAll(ctx context.Context) []web.HakAksesListResponse {
	tx, err := service.Db.Begin()
	defer helpers.CommitOrRollback(tx, err)

	aksess := service.HakAksesRepository.FindAll(ctx, tx)

	return web.ToHakAksesListResponses(aksess)
}
