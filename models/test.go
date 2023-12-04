package models

type Test struct {
	Id   int    `gorm:"primaryKey" json:"id" form:"id"`
	Name string `json:"name" form:"name"`
	Addr string `json:"addr" form:"addr"`
	Star string `json:"star" form:"star"`
}

// type Test struct {
// 	Id   int    `json:"id" form:"id"`
// 	Name string `json:"name" form:"name"`
// 	Addr string `json:"addr" form:"addr"`
// 	Star string `json:"star" form:"star"`
// }
