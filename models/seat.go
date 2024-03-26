package models

type UsedBy uint8

const (
	UsedByNop     UsedBy = 0 // 无
	UsedByIn      UsedBy = 1 // 占用
	UsedByReserve UsedBy = 2 // 预约

)

// Seat  座位表
type Seat struct {
	Model
	PlaceId int32  `json:"place_id"`
	Number  string `json:"number"`  // 座位号 楼层-编号
	UsedBy  UsedBy `json:"used_by"` // 业务状态
}
