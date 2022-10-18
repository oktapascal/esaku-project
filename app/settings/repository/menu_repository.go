package repository

import (
	"esaku-project/app/settings/models/domain"
)

type MenuRepository interface {
	Save(menu domain.Menu)
	Update(menu domain.Menu)
	FindById(KlpMenu string) (domain.Menu, error)
}
