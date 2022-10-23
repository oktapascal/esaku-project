package web

import (
	"esaku-project/app/settings/models/domain"
)

type dataMenuSub struct {
	KodeMenu string        `json:"kode_menu" validate:"required,min=1,max=10"`
	NamaMenu string        `json:"nama_menu" validate:"required,min=1,max=100"`
	KodeForm string        `json:"kode_form" validate:"max=10"`
	Level    int           `json:"level" validate:"number"`
	SubMenu  []dataMenuSub `json:"data_menu" validate:"dive"`
}

type dataMenuMain struct {
	KodeMenu string        `json:"kode_menu" validate:"required,min=1,max=10"`
	NamaMenu string        `json:"nama_menu" validate:"required,min=1,max=100"`
	KodeForm string        `json:"kode_form" validate:"max=10"`
	Level    int           `json:"level" validate:"number"`
	SubMenu  []dataMenuSub `json:"data_menu" validate:"dive"`
}

type MenuSaveRequest struct {
	KodeKlpMenu string         `json:"kode_klp_menu" validate:"required,min=1,max=20"`
	ListMenu    []dataMenuMain `json:"data_menu" validate:"required,min=1,dive"`
}

type MenuPosition struct {
	KodeMenu string
	NamaMenu string
	Level    int
	Index    int
}

func ToDomainKelompokMenu(menu MenuSaveRequest) domain.KelompokMenu {
	return domain.KelompokMenu{
		KodeKlp: menu.KodeKlpMenu,
	}
}

func ToDomainForm(kodeForm string) domain.Form {
	return domain.Form{
		KodeForm: kodeForm,
	}
}

func ToDomainMenu(menu MenuPosition, klpMenu domain.KelompokMenu, form domain.Form) domain.Menu {
	return domain.Menu{
		KodeMenu:     menu.KodeMenu,
		NamaMenu:     menu.NamaMenu,
		Level:        menu.Level,
		Index:        menu.Index,
		KelompokMenu: klpMenu,
		Form:         form,
	}
}

func ToDomainMenuRequests(menu MenuSaveRequest) []domain.Menu {
	var kelompokMenu domain.KelompokMenu
	var form domain.Form
	var menuPosition MenuPosition

	var menus []domain.Menu

	kelompokMenu = ToDomainKelompokMenu(menu)

	i := 0
	for _, menuMain := range menu.ListMenu {
		form = ToDomainForm(menuMain.KodeForm)

		menuPosition.KodeMenu = menuMain.KodeMenu
		menuPosition.NamaMenu = menuMain.NamaMenu
		menuPosition.Level = menuMain.Level
		menuPosition.Index = i

		menus = append(menus, ToDomainMenu(menuPosition, kelompokMenu, form))

		if len(menuMain.SubMenu) > 0 {
			for _, menuSub1 := range menuMain.SubMenu {
				i++
				form = ToDomainForm(menuSub1.KodeForm)

				menuPosition.KodeMenu = menuSub1.KodeMenu
				menuPosition.NamaMenu = menuSub1.NamaMenu
				menuPosition.Level = menuSub1.Level
				menuPosition.Index = i

				menus = append(menus, ToDomainMenu(menuPosition, kelompokMenu, form))
				if len(menuSub1.SubMenu) > 0 {
					for _, menuSub2 := range menuSub1.SubMenu {
						i++
						form = ToDomainForm(menuSub2.KodeForm)

						menuPosition.KodeMenu = menuSub2.KodeMenu
						menuPosition.NamaMenu = menuSub2.NamaMenu
						menuPosition.Level = menuSub2.Level
						menuPosition.Index = i

						menus = append(menus, ToDomainMenu(menuPosition, kelompokMenu, form))
					}
				}
			}
		}

		i++
	}

	return menus
}
