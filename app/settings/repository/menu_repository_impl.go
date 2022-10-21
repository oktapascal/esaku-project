package repository

import (
	"context"
	"database/sql"
	"esaku-project/app/settings/models/domain"
	"esaku-project/helpers"
	"fmt"
	"strings"
)

type MenuRepositoryImpl struct {
}

func NewMenuRepositoryImpl() *MenuRepositoryImpl {
	return &MenuRepositoryImpl{}
}

func (repository *MenuRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, menu []domain.Menu) {
	valueStrings := make([]string, 0, len(menu))
	valueArgs := make([]string, 0, len(menu)*6)

	i := 0
	for _, menu := range menu {
		valueStrings = append(valueStrings, fmt.Sprintf("(@p%d, @p%d, @p%d, @p%d, @p%d, @p%d)", i*6+1, i*6+2, i*6+3, i*6+4, i*6+5, i*6+6))
		valueArgs = append(valueArgs, menu.KodeKlp)
		valueArgs = append(valueArgs, menu.KodeMenu)
		valueArgs = append(valueArgs, menu.NamaMenu)
		valueArgs = append(valueArgs, menu.Level)
		valueArgs = append(valueArgs, menu.Index)
		valueArgs = append(valueArgs, menu.KodeForm)
		i++
	}

	SQL := fmt.Sprintf(`insert into menu (kode_klp, kode_menu, nama_menu, level_menu, rowindex, kode_form values %s)`, strings.Join(valueStrings, ","))
	fmt.Println(SQL)
}

func (repository *MenuRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, kodeKlp string) {
	SQL := "delete from menu where kode_klp = @p1"

	_, err := tx.ExecContext(ctx, SQL, kodeKlp)
	helpers.PanicIfError(err)
}

func (repository *MenuRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, kodeKlp string) []domain.Menu {
	SQL := `select a.kode_klp, a.kode_menu, a.nama nama_menu, a.level_menu, a.rowindex, a.kode_form,
	isnull(b.nama, '-') nama_klp_menu, isnull(c.nama_form, '-') nama_form
	from menu a
	left join menu_klp b on a.kode_klp=b.kode_klp
	left join m_form c on a.kode_form=c.kode_form
	where a.kode_klp = @p1
	order by a.rowindex asc`

	rows, err := tx.QueryContext(ctx, SQL, kodeKlp)
	helpers.PanicIfError(err)
	//noinspection GoUnhandledErrorResult
	defer rows.Close()

	var menus []domain.Menu

	for rows.Next() {
		menu := domain.Menu{}

		err := rows.Scan(&menu.KelompokMenu.KodeKlp, &menu.KodeMenu, &menu.NamaMenu, &menu.Level, &menu.Index,
			&menu.Form.KodeForm, &menu.KelompokMenu.Nama, &menu.Form.Nama)
		helpers.PanicIfError(err)

		menus = append(menus, menu)
	}

	return menus
}
