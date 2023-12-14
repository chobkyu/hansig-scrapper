package models

type Hansic struct {
	Id         int    `gorm:"primaryKey" json:"id" form:"id"`
	Name       string `json:"name" form:"name"`
	Addr       string `json:"addr" form:"addr"`
	GoogleStar string `json:"googleStar" form:"googleStar"`
	LocationId int    `json:"location" form:"location"`
}

// type Test struct {
// 	Id   int    `json:"id" form:"id"`
// 	Name string `json:"name" form:"name"`
// 	Addr string `json:"addr" form:"addr"`
// 	Star string `json:"star" form:"star"`
// }
