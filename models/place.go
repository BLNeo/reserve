package models

// Place 场所表
type Place struct {
	Model
	Name     string `json:"name"`      // 名称
	ShowName string `json:"show_name"` //  展示名称
	City     string `json:"city"`
	Address  string `json:"address"` // 地址
}
