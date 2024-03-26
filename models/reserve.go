package models

import "time"

type ReserveStatus uint8

const (
	ReserveStatusNop     ReserveStatus = 0 // 未到
	ReserveStatusIn      ReserveStatus = 1 // 已到
	ReserveStatusCancel  ReserveStatus = 2 // 取消
	ReserveStatusExpires ReserveStatus = 3 // 过期 （开始时间15分钟未签到）
	ReserveStatusAddTime ReserveStatus = 4 // 延长使用
)

// Reserve 预约表
type Reserve struct {
	Model
	UserId  int32         `json:"user_id"`
	SeatId  int32         `json:"seat_id"`
	PlaceId int32         `json:"place_id"`
	Begin   time.Time     `json:"begin"`   // 预约的开始时间
	End     time.Time     `json:"end"`     // 预约的结束时间
	InTime  time.Time     `json:"in_time"` // 签到时间
	Status  ReserveStatus `json:"status"`  // 预约状态
}
