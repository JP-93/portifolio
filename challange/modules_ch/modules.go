package modules_ch

type Revenue struct {
	ID        uint     `gorm:"primary key;autoIncremente" json:"ID"`
	Descricao *string  `json:"descricao"`
	Valor     *float64 `json:"valor"`
	Data_into *string  `json:"data_into"`
}

type Expense struct {
	ID        uint     `gorm:"primary key;autoIncremente" json:"ID"`
	Descricao *string  `json:"descricao"`
	Valor     *float64 `json:"valor"`
	Data_into *string  `json:"data_into"`
	Categoria string   `grom:"default 'outros'" json:"categoria"`
}
