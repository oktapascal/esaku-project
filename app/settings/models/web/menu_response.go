package web

import (
	"esaku-project/app/settings/models/domain"
	"strconv"
)

type listMenuSub struct {
	KodeMenu string        `json:"kode_menu"`
	NamaMenu string        `json:"nama_menu"`
	Program  string        `json:"program"`
	Index    uint64        `json:"row_index"`
	SubMenu  []listMenuSub `json:"sub_menu"`
}

type listMenuMain struct {
	KodeMenu string        `json:"kode_menu"`
	NamaMenu string        `json:"nama_menu"`
	Program  string        `json:"program"`
	Index    uint64        `json:"row_index"`
	SubMenu  []listMenuSub `json:"sub_menu"`
}

type MenuResponse struct {
	KodeKlpMenu string         `json:"kode_klp_menu"`
	NamaKlpMenu string         `json:"nama_klp_menu"`
	ListMenu    []listMenuMain `json:"data_menu"`
}

func toListMenuSub(menu domain.Menu) listMenuSub {
	idx, _ := strconv.ParseUint(menu.Index, 10, 64)
	return listMenuSub{
		KodeMenu: menu.KodeMenu,
		NamaMenu: menu.NamaMenu,
		Program:  menu.Program,
		Index:    idx,
		SubMenu:  nil,
	}
}

func toListMenuMain(menu domain.Menu) listMenuMain {
	idx, _ := strconv.ParseUint(menu.Index, 10, 64)
	return listMenuMain{
		KodeMenu: menu.KodeMenu,
		NamaMenu: menu.NamaMenu,
		Program:  menu.Program,
		Index:    idx,
		SubMenu:  nil,
	}
}

func toListMenuResponse(menu []domain.Menu) []listMenuMain {
	var listMenuResponses []listMenuMain

	var level uint64
	i := -1
	j := -1
	for _, menu := range menu {
		level, _ = strconv.ParseUint(menu.Level, 10, 64)
		if level == 0 {
			listMenuResponses = append(listMenuResponses, toListMenuMain(menu))

			i++
		}

		if level == 1 {
			listMenuResponses[i].SubMenu = append(listMenuResponses[i].SubMenu, toListMenuSub(menu))
			j++
		}

		if level == 2 {
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
