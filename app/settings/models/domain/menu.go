package domain

type SubMenu3 struct {
	Nama    string
	Program string
	SubMenu []SubMenu2
}

type SubMenu2 struct {
	Nama    string
	Program string
}

type SubMenu1 struct {
	Nama    string
	Program string
	SubMenu []SubMenu2
}

type Menu struct {
	KodeKlpMenu string
	SubMenu     []SubMenu1
}
