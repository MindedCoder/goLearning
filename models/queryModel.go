package models

type QueryModel struct {
	Where   string `form:"where"`
	Include string 					  `form:"include"`
	Skip    int 						`form:"skip"`
	Limit   int								`form:"limit"`
	Count   int								`form:"count"`
	Order   string						`form:"order"`
}
