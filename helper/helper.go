package helper

type Pagination struct {
	Totalpage     int
	Prevpage      int
	Nextpage      int
	Totalrows     int
	TotalNextpage int
	Totalprevpage int
}
type LimitPage struct {
	PName string `form:"pname"`
	Limit int    `form:"limit"`
	Page  int    `form:"page"`
}

type UserFilter struct {
	Fname string `form:"fname"`
	Lname string `form:"lname"`
	Email string `form:"email"`
	Limit int    `form:"limit"`
	Page  int    `form:"page"`
}
