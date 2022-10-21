package controllers

import (
	"esaku-project/app/settings/services"
	"net/http"
)

type MenuControllerImpl struct {
	MenuService services.MenuService
}

func NewMenuControllerImpl(menuService services.MenuService) *MenuControllerImpl {
	return &MenuControllerImpl{MenuService: menuService}
}

func (controller *MenuControllerImpl) Save(writer http.ResponseWriter, request *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (controller *MenuControllerImpl) Delete(writer http.ResponseWriter, request *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (controller *MenuControllerImpl) FindById(writer http.ResponseWriter, request *http.Request) {
	//TODO implement me
	panic("implement me")
}
