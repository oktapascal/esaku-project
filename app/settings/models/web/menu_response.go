package web

import "esaku-project/app/settings/models/domain"

type listMenu struct {
	KodeMenu string `json:"kode_menu"`
	NamaMenu string `json:"nama_menu"`
	Program  string `json:"program"`
	Level    string `json:"level_menu"`
	Index    string `json:"row_index"`
}

type MenuResponse struct {
	KodeKlpMenu string     `json:"kode_klp_menu"`
	NamaKlpMenu string     `json:"nama_klp_menu"`
	ListMenu    []listMenu `json:"data_menu"`
}

func toListMenu(menu domain.Menu) listMenu {
	return listMenu{
		KodeMenu: menu.KodeMenu,
		NamaMenu: menu.NamaMenu,
		Program:  menu.Program,
		Level:    menu.Level,
		Index:    menu.Index,
	}
}

func toListMenuResponse(menu []domain.Menu) []listMenu {
	var listMenuResponses []listMenu

	for _, menu := range menu {
		listMenuResponses = append(listMenuResponses, toListMenu(menu))
	}

	return listMenuResponses
}

func ToMenuResponse(klpMenu KelompokMenuResponse, menus []domain.Menu) MenuResponse {

	return MenuResponse{
		KodeKlpMenu: klpMenu.KodeKlp,
		NamaKlpMenu: klpMenu.Nama,
		ListMenu:    nil,
	}
}
