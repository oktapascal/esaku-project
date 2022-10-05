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

type FormServiceImpl struct {
	FormRepository repository.FormRepository
	Db             *sql.DB
	Validate       *validator.Validate
}

func NewFormServiceImpl(formRepository repository.FormRepository, db *sql.DB, validate *validator.Validate) *FormServiceImpl {
	return &FormServiceImpl{
		FormRepository: formRepository,
		Db:             db,
		Validate:       validate,
	}
}

func (service *FormServiceImpl) Save(ctx context.Context, request web.FormSaveRequest) web.FormResponse {
	err := service.Validate.Struct(request)

	helpers.PanicIfError(err)

	tx, err := service.Db.Begin()

	defer helpers.CommitOrRollback(tx, err)

	form := domain.Form{
		KodeForm: request.KodeForm,
		Nama:     request.Nama,
		Program:  request.Program,
	}

	form = service.FormRepository.Save(ctx, tx, form)

	return web.ToFormResponse(form)
}

func (service *FormServiceImpl) Update(ctx context.Context, request web.FormSaveRequest) web.FormResponse {
	err := service.Validate.Struct(request)

	helpers.PanicIfError(err)

	tx, err := service.Db.Begin()

	defer helpers.CommitOrRollback(tx, err)

	form, err := service.FormRepository.FindById(ctx, tx, request.KodeForm)

	if err != nil {
		panic(exceptions.NewErrorNotFound(err.Error()))
	}

	form.Nama = request.Nama
	form.Program = request.Program

	form = service.FormRepository.Update(ctx, tx, form)

	return web.ToFormResponse(form)
}

func (service *FormServiceImpl) Delete(ctx context.Context, kodeForm string) {
	tx, err := service.Db.Begin()

	defer helpers.CommitOrRollback(tx, err)

	form, err := service.FormRepository.FindById(ctx, tx, kodeForm)

	if err != nil {
		panic(exceptions.NewErrorNotFound(err.Error()))
	}

	service.FormRepository.Delete(ctx, tx, form)
}

func (service *FormServiceImpl) FindById(ctx context.Context, kodeForm string) web.FormResponse {
	tx, err := service.Db.Begin()

	defer helpers.CommitOrRollback(tx, err)

	form, err := service.FormRepository.FindById(ctx, tx, kodeForm)

	if err != nil {
		panic(exceptions.NewErrorNotFound(err.Error()))
	}

	return web.ToFormResponse(form)
}

func (service *FormServiceImpl) FindAll(ctx context.Context) []web.FormResponse {
	tx, err := service.Db.Begin()

	defer helpers.CommitOrRollback(tx, err)

	forms := service.FormRepository.FindAll(ctx, tx)

	return web.ToFormResponses(forms)
}
