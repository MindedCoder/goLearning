package models

type QueryModel struct {
	Where   string `form:"where"`
	Include string 					  `form:"include"`
	Skip    int 						`form:"skip"`
	Limit   int								`form:"limit"`
	Order   string						`form:"order"`
}
