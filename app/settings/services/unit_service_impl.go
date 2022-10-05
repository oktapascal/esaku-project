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

type UnitServiceImpl struct {
	UnitRepository repository.UnitRepository
	Db             *sql.DB
	Validate       *validator.Validate
}

func NewUnitServiceImpl(unitRepository repository.UnitRepository, db *sql.DB, validate *validator.Validate) *UnitServiceImpl {
	return &UnitServiceImpl{
		UnitRepository: unitRepository,
		Db:             db,
		Validate:       validate,
	}
}

func (service *UnitServiceImpl) Save(ctx context.Context, request web.UnitSaveRequest) web.UnitResponse {
	err := service.Validate.Struct(request)

	helpers.PanicIfError(err)

	tx, err := service.Db.Begin()

	defer helpers.CommitOrRollback(tx, err)

	kodeLokasi := "99"
	unit := domain.Unit{
		KodeUnit:   request.KodeUnit,
		KodeLokasi: kodeLokasi,
		Nama:       request.Nama,
		FlagAktif:  request.FlagAktif,
	}

	unit = service.UnitRepository.Save(ctx, tx, unit)

	return web.ToUnitResponse(unit)
}

func (service *UnitServiceImpl) Update(ctx context.Context, request web.UnitUpdateRequest) web.UnitResponse {
	err := service.Validate.Struct(request)

	helpers.PanicIfError(err)

	tx, err := service.Db.Begin()

	defer helpers.CommitOrRollback(tx, err)

	kodeLokasi := "99"
	unit, err := service.UnitRepository.FindById(ctx, tx, request.KodeUnit, kodeLokasi)

	if err != nil {
		panic(exceptions.NewErrorNotFound(err.Error()))
	}

	unit.Nama = request.Nama
	unit.FlagAktif = request.FlagAktif
	unit.KodeLokasi = kodeLokasi

	unit = service.UnitRepository.Update(ctx, tx, unit)

	return web.ToUnitResponse(unit)
}

func (service *UnitServiceImpl) Delete(ctx context.Context, kodeUnit string) {
	tx, err := service.Db.Begin()

	defer helpers.CommitOrRollback(tx, err)

	kodeLokasi := "99"
	unit, err := service.UnitRepository.FindById(ctx, tx, kodeUnit, kodeLokasi)

	if err != nil {
		panic(exceptions.NewErrorNotFound(err.Error()))
	}

	service.UnitRepository.Delete(ctx, tx, unit)
}

func (service *UnitServiceImpl) FindById(ctx context.Context, kodeUnit string) web.UnitResponse {
	tx, err := service.Db.Begin()

	defer helpers.CommitOrRollback(tx, err)

	kodeLokasi := "99"
	unit, err := service.UnitRepository.FindById(ctx, tx, kodeUnit, kodeLokasi)

	if err != nil {
		panic(exceptions.NewErrorNotFound(err.Error()))
	}

	return web.ToUnitResponse(unit)
}

func (service *UnitServiceImpl) FindAll(ctx context.Context) []web.UnitResponse {
	tx, err := service.Db.Begin()

	defer helpers.CommitOrRollback(tx, err)

	kodeLokasi := "99"
	units := service.UnitRepository.FindAll(ctx, tx, kodeLokasi)

	return web.ToUnitResponses(units)
}
