package repository

import (
	"esaku-project/app/settings/models/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type MenuRepositoryImpl struct {
	Collection *mongo.Collection
}

func NewMenuRepositoryImpl(database *mongo.Database) *MenuRepositoryImpl {
	return &MenuRepositoryImpl{
		Collection: database.Collection("menus"),
	}
}

func (repository *MenuRepositoryImpl) Save(menu domain.Menu) {
	//TODO implement me
	panic("implement me")
}

func (repository *MenuRepositoryImpl) Update(menu domain.Menu) {
	//TODO implement me
	panic("implement me")
}

func (repository *MenuRepositoryImpl) FindById(KlpMenu string) (domain.Menu, error) {
	//TODO implement me
	panic("implement me")
}
