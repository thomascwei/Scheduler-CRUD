package internal

import (
	"schedule/pkg/db"
	"time"
)

type Config struct {
	Host     string `mapstructure:"DB_HOST"`
	Port     string `mapstructure:"DB_PORT"`
	User     string `mapstructure:"DB_USER"`
	Password string `mapstructure:"DB_PASSWORD"`
}

type RequestId struct {
	Id int32 `json:"id"`
}
type RequestScheduleOne struct {
	ID int32 `json:"id"`
	// template type
	TimeTypeID  int32 `json:"time_type_id" binding:"required,oneof=0 1 2 3"`
	IntervalDay int32 `json:"interval_day" binding:"required"`
	// repeat in day
	IntervalSeconds int32  `json:"interval_seconds" binding:"required"`
	AtTime          string `json:"at_time" binding:"required"`
	StartTime       string `json:"start_time" binding:"required"`
	EndTime         string `json:"end_time" binding:"required"`
	CommandID       int32  `json:"command_id" binding:"required"`
	// template name
	Name          string    `json:"name" binding:"required"`
	StartDate     time.Time `json:"start_date" binding:"required"`
	EndDate       time.Time `json:"end_date" binding:"required"`
	Enable        *bool     `json:"enable" binding:"required"`
	Repeat        *bool     `json:"repeat" binding:"required"`
	RepeatWeekday []string  `json:"repeat_weekday" binding:"required"`
	RepeatDay     []string  `json:"repeat_day" binding:"required"`
	RepeatMonth   []string  `json:"repeat_month" binding:"required"`
}

type ResponseScheduleOne struct {
	db.Schedule
	RepeatWeekday []string `json:"repeat_weekday" binding:"required"`
	RepeatDay     []string `json:"repeat_day" binding:"required"`
	RepeatMonth   []string `json:"repeat_month" binding:"required"`
}

type SubSchedules struct {
	RepeatWeekday []db.ListRepeatWeekdaysRow `json:"repeat_weekday"`
	RepeatMonth   []db.ListRepeatMonthsRow   `json:"repeat_month"`
	RepeatDay     []db.ListRepeatDaysRow     `json:"repeat_day"`
}
