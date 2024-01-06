package models

type Registration struct {
	Id      uint
	UserId  uint `binding:"required"`
	EventId uint `binding:"required"`
}
