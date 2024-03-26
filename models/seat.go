package models

type SeatStatus uint8

const (
	SeatStatusNop     SeatStatus = 0 // 无
	SeatStatusIn      SeatStatus = 1 // 占用中 (已完成签到)
	SeatStatusReserve SeatStatus = 2 // 预约
	SeatStatusLock    SeatStatus = 3 // 锁定
	SeatStatusRepair  SeatStatus = 4 // 维修
)

// Seat  座位表
type Seat struct {
	Model
	PlaceId int32      `json:"place_id"`
	Number  string     `json:"number"` // 座位号 楼层-编号
	Status  SeatStatus `json:"status"` // Status 业务相关状态
}
