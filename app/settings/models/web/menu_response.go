package web

import (
	"esaku-project/app/settings/models/domain"
)

type listMenuSub struct {
	KodeMenu string        `json:"kode_menu"`
	NamaMenu string        `json:"nama_menu"`
	Program  string        `json:"program"`
	Index    int           `json:"row_index"`
	SubMenu  []listMenuSub `json:"sub_menu"`
}

type listMenuMain struct {
	KodeMenu string        `json:"kode_menu"`
	NamaMenu string        `json:"nama_menu"`
	Program  string        `json:"program"`
	Index    int           `json:"row_index"`
	SubMenu  []listMenuSub `json:"sub_menu"`
}

type MenuResponse struct {
	KodeKlpMenu string         `json:"kode_klp_menu"`
	NamaKlpMenu string         `json:"nama_klp_menu"`
	ListMenu    []listMenuMain `json:"data_menu"`
}

func toListMenuSub(menu domain.Menu) listMenuSub {
	return listMenuSub{
		KodeMenu: menu.KodeMenu,
		NamaMenu: menu.NamaMenu,
		Program:  menu.Program,
		Index:    menu.Index,
		SubMenu:  nil,
	}
}

func toListMenuMain(menu domain.Menu) listMenuMain {
	return listMenuMain{
		KodeMenu: menu.KodeMenu,
		NamaMenu: menu.NamaMenu,
		Program:  menu.Program,
		Index:    menu.Index,
		SubMenu:  nil,
	}
}

func toListMenuResponse(menu []domain.Menu) []listMenuMain {
	var listMenuResponses []listMenuMain

	i := -1
	j := -1
	for _, menu := range menu {
		if menu.Level == 0 {
			listMenuResponses = append(listMenuResponses, toListMenuMain(menu))

			i++
		}

		if menu.Level == 1 {
			listMenuResponses[i].SubMenu = append(listMenuResponses[i].SubMenu, toListMenuSub(menu))
			j++
		}

		if menu.Level == 2 {
			listMenuResponses[i].SubMenu[j].SubMenu = append(listMenuResponses[i].SubMenu[j].SubMenu, toListMenuSub(menu))
		}
	}

	return listMenuResponses
}

func ToMenuResponse(klpMenu domain.KelompokMenu, menus []domain.Menu) MenuResponse {
	listMenu := toListMenuResponse(menus)
	return MenuResponse{
		KodeKlpMenu: klpMenu.KodeKlp,
		NamaKlpMenu: klpMenu.Nama,
		ListMenu:    listMenu,
	}
}
