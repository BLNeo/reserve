package book

type AddBookRequest struct {
	Name     string  `json:"name" binding:"required"`
	StoreNum int     `json:"store_num" `
	Price    float64 `json:"price" binding:"required"`
}

type ListBookRequest struct {
	UserId int64 `json:"user_id"`
}

type ListBookRespond struct {
	Name     string  `json:"name" `
	StoreNum int     `json:"store_num"`
	Price    float64 `json:"price" `
}
