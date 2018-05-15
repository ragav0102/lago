package models

import (
	"github.com/jinzhu/gorm"
)

type WorkOrder struct {
	gorm.Model
	ExternalId uint
	OrderType  string
	ContactId  uint
	Address    string
	ScheduleId uint
	AgentId    uint
}
