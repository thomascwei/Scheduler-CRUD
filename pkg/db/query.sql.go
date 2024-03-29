// Code generated by sqlc. DO NOT EDIT.
// source: query.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const CreateCommand = `-- name: CreateCommand :execresult
INSERT INTO commands(command)
VALUES (?)
`

func (q *Queries) CreateCommand(ctx context.Context, command string) (sql.Result, error) {
	return q.exec(ctx, q.createCommandStmt, CreateCommand, command)
}

const CreateRepeatDays = `-- name: CreateRepeatDays :exec
INSERT INTO repeat_day(schedule_id, day)
VALUES (?, ?)
`

type CreateRepeatDaysParams struct {
	ScheduleID int32  `json:"schedule_id"`
	Day        string `json:"day"`
}

func (q *Queries) CreateRepeatDays(ctx context.Context, arg CreateRepeatDaysParams) error {
	_, err := q.exec(ctx, q.createRepeatDaysStmt, CreateRepeatDays, arg.ScheduleID, arg.Day)
	return err
}

const CreateRepeatMonth = `-- name: CreateRepeatMonth :exec
INSERT INTO repeat_month(schedule_id, month)
VALUES (?, ?)
`

type CreateRepeatMonthParams struct {
	ScheduleID int32  `json:"schedule_id"`
	Month      string `json:"month"`
}

func (q *Queries) CreateRepeatMonth(ctx context.Context, arg CreateRepeatMonthParams) error {
	_, err := q.exec(ctx, q.createRepeatMonthStmt, CreateRepeatMonth, arg.ScheduleID, arg.Month)
	return err
}

const CreateRepeatWeekdays = `-- name: CreateRepeatWeekdays :exec
INSERT INTO repeat_weekday(schedule_id, weekday)
VALUES (?, ?)
`

type CreateRepeatWeekdaysParams struct {
	ScheduleID int32  `json:"schedule_id"`
	Weekday    string `json:"weekday"`
}

func (q *Queries) CreateRepeatWeekdays(ctx context.Context, arg CreateRepeatWeekdaysParams) error {
	_, err := q.exec(ctx, q.createRepeatWeekdaysStmt, CreateRepeatWeekdays, arg.ScheduleID, arg.Weekday)
	return err
}

const CreateSchedule = `-- name: CreateSchedule :execresult
INSERT INTO schedules (time_type_id, interval_day, interval_seconds, at_time, start_time, end_time,
                       command_id, name, start_date, end_date, enable, ` + "`" + `repeat` + "`" + `)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
`

type CreateScheduleParams struct {
	TimeTypeID      int32     `json:"time_type_id"`
	IntervalDay     int32     `json:"interval_day"`
	IntervalSeconds int32     `json:"interval_seconds"`
	AtTime          string    `json:"at_time"`
	StartTime       string    `json:"start_time"`
	EndTime         string    `json:"end_time"`
	CommandID       int32     `json:"command_id"`
	Name            string    `json:"name"`
	StartDate       time.Time `json:"start_date"`
	EndDate         time.Time `json:"end_date"`
	Enable          bool      `json:"enable"`
	Repeat          bool      `json:"repeat"`
}

func (q *Queries) CreateSchedule(ctx context.Context, arg CreateScheduleParams) (sql.Result, error) {
	return q.exec(ctx, q.createScheduleStmt, CreateSchedule,
		arg.TimeTypeID,
		arg.IntervalDay,
		arg.IntervalSeconds,
		arg.AtTime,
		arg.StartTime,
		arg.EndTime,
		arg.CommandID,
		arg.Name,
		arg.StartDate,
		arg.EndDate,
		arg.Enable,
		arg.Repeat,
	)
}

const DeleteCommand = `-- name: DeleteCommand :exec
DELETE
FROM commands
WHERE id = ?
`

func (q *Queries) DeleteCommand(ctx context.Context, id int32) error {
	_, err := q.exec(ctx, q.deleteCommandStmt, DeleteCommand, id)
	return err
}

const DeleteRepeatDay = `-- name: DeleteRepeatDay :exec
DELETE
FROM repeat_day
WHERE schedule_id = ?
`

func (q *Queries) DeleteRepeatDay(ctx context.Context, scheduleID int32) error {
	_, err := q.exec(ctx, q.deleteRepeatDayStmt, DeleteRepeatDay, scheduleID)
	return err
}

const DeleteRepeatMonth = `-- name: DeleteRepeatMonth :exec
DELETE
FROM repeat_month
WHERE schedule_id = ?
`

func (q *Queries) DeleteRepeatMonth(ctx context.Context, scheduleID int32) error {
	_, err := q.exec(ctx, q.deleteRepeatMonthStmt, DeleteRepeatMonth, scheduleID)
	return err
}

const DeleteRepeatWeekday = `-- name: DeleteRepeatWeekday :exec
DELETE
FROM repeat_weekday
WHERE schedule_id = ?
`

func (q *Queries) DeleteRepeatWeekday(ctx context.Context, scheduleID int32) error {
	_, err := q.exec(ctx, q.deleteRepeatWeekdayStmt, DeleteRepeatWeekday, scheduleID)
	return err
}

const DeleteSchedule = `-- name: DeleteSchedule :exec
DELETE
FROM schedules
WHERE id = ?
`

func (q *Queries) DeleteSchedule(ctx context.Context, id int32) error {
	_, err := q.exec(ctx, q.deleteScheduleStmt, DeleteSchedule, id)
	return err
}

const GetCommands = `-- name: GetCommands :many
SELECT id, command, create_time
FROM commands
`

func (q *Queries) GetCommands(ctx context.Context) ([]Command, error) {
	rows, err := q.query(ctx, q.getCommandsStmt, GetCommands)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Command
	for rows.Next() {
		var i Command
		if err := rows.Scan(&i.ID, &i.Command, &i.CreateTime); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const GetDays = `-- name: GetDays :many
SELECT id, schedule_id, day, create_time
FROM repeat_day
WHERE schedule_id = ?
`

func (q *Queries) GetDays(ctx context.Context, scheduleID int32) ([]RepeatDay, error) {
	rows, err := q.query(ctx, q.getDaysStmt, GetDays, scheduleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []RepeatDay
	for rows.Next() {
		var i RepeatDay
		if err := rows.Scan(
			&i.ID,
			&i.ScheduleID,
			&i.Day,
			&i.CreateTime,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const GetMonths = `-- name: GetMonths :many
SELECT id, schedule_id, month, create_time
FROM repeat_month
WHERE schedule_id = ?
`

func (q *Queries) GetMonths(ctx context.Context, scheduleID int32) ([]RepeatMonth, error) {
	rows, err := q.query(ctx, q.getMonthsStmt, GetMonths, scheduleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []RepeatMonth
	for rows.Next() {
		var i RepeatMonth
		if err := rows.Scan(
			&i.ID,
			&i.ScheduleID,
			&i.Month,
			&i.CreateTime,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const GetSchedule = `-- name: GetSchedule :one
SELECT id, time_type_id, interval_day, interval_seconds, at_time, start_time, end_time, command_id, name, start_date, end_date, enable, ` + "`" + `repeat` + "`" + `, create_time
FROM schedules
WHERE id = ?
LIMIT 1
`

func (q *Queries) GetSchedule(ctx context.Context, id int32) (Schedule, error) {
	row := q.queryRow(ctx, q.getScheduleStmt, GetSchedule, id)
	var i Schedule
	err := row.Scan(
		&i.ID,
		&i.TimeTypeID,
		&i.IntervalDay,
		&i.IntervalSeconds,
		&i.AtTime,
		&i.StartTime,
		&i.EndTime,
		&i.CommandID,
		&i.Name,
		&i.StartDate,
		&i.EndDate,
		&i.Enable,
		&i.Repeat,
		&i.CreateTime,
	)
	return i, err
}

const GetWeekdays = `-- name: GetWeekdays :many
SELECT id, schedule_id, weekday, create_time
FROM repeat_weekday
WHERE schedule_id = ?
`

func (q *Queries) GetWeekdays(ctx context.Context, scheduleID int32) ([]RepeatWeekday, error) {
	rows, err := q.query(ctx, q.getWeekdaysStmt, GetWeekdays, scheduleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []RepeatWeekday
	for rows.Next() {
		var i RepeatWeekday
		if err := rows.Scan(
			&i.ID,
			&i.ScheduleID,
			&i.Weekday,
			&i.CreateTime,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const ListRepeatDays = `-- name: ListRepeatDays :many
SELECT schedule_id, day
FROM repeat_day
ORDER BY id
`

type ListRepeatDaysRow struct {
	ScheduleID int32  `json:"schedule_id"`
	Day        string `json:"day"`
}

func (q *Queries) ListRepeatDays(ctx context.Context) ([]ListRepeatDaysRow, error) {
	rows, err := q.query(ctx, q.listRepeatDaysStmt, ListRepeatDays)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListRepeatDaysRow
	for rows.Next() {
		var i ListRepeatDaysRow
		if err := rows.Scan(&i.ScheduleID, &i.Day); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const ListRepeatMonths = `-- name: ListRepeatMonths :many
SELECT schedule_id, month
FROM repeat_month
ORDER BY id
`

type ListRepeatMonthsRow struct {
	ScheduleID int32  `json:"schedule_id"`
	Month      string `json:"month"`
}

func (q *Queries) ListRepeatMonths(ctx context.Context) ([]ListRepeatMonthsRow, error) {
	rows, err := q.query(ctx, q.listRepeatMonthsStmt, ListRepeatMonths)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListRepeatMonthsRow
	for rows.Next() {
		var i ListRepeatMonthsRow
		if err := rows.Scan(&i.ScheduleID, &i.Month); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const ListRepeatWeekdays = `-- name: ListRepeatWeekdays :many
SELECT schedule_id, weekday
FROM repeat_weekday
ORDER BY id
`

type ListRepeatWeekdaysRow struct {
	ScheduleID int32  `json:"schedule_id"`
	Weekday    string `json:"weekday"`
}

func (q *Queries) ListRepeatWeekdays(ctx context.Context) ([]ListRepeatWeekdaysRow, error) {
	rows, err := q.query(ctx, q.listRepeatWeekdaysStmt, ListRepeatWeekdays)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListRepeatWeekdaysRow
	for rows.Next() {
		var i ListRepeatWeekdaysRow
		if err := rows.Scan(&i.ScheduleID, &i.Weekday); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const ListSchedule = `-- name: ListSchedule :many
SELECT id, time_type_id, interval_day, interval_seconds, at_time, start_time, end_time, command_id, name, start_date, end_date, enable, ` + "`" + `repeat` + "`" + `, create_time
FROM schedules
ORDER BY id
`

func (q *Queries) ListSchedule(ctx context.Context) ([]Schedule, error) {
	rows, err := q.query(ctx, q.listScheduleStmt, ListSchedule)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Schedule
	for rows.Next() {
		var i Schedule
		if err := rows.Scan(
			&i.ID,
			&i.TimeTypeID,
			&i.IntervalDay,
			&i.IntervalSeconds,
			&i.AtTime,
			&i.StartTime,
			&i.EndTime,
			&i.CommandID,
			&i.Name,
			&i.StartDate,
			&i.EndDate,
			&i.Enable,
			&i.Repeat,
			&i.CreateTime,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const UpdateCommand = `-- name: UpdateCommand :exec
UPDATE commands
set command=?
WHERE id = ?
`

type UpdateCommandParams struct {
	Command string `json:"command"`
	ID      int32  `json:"id"`
}

func (q *Queries) UpdateCommand(ctx context.Context, arg UpdateCommandParams) error {
	_, err := q.exec(ctx, q.updateCommandStmt, UpdateCommand, arg.Command, arg.ID)
	return err
}

const UpdateSchedule = `-- name: UpdateSchedule :exec
UPDATE schedules
SET time_type_id=?,
    interval_day=?,
    interval_seconds=?,
    at_time=?,
    start_time=?,
    end_time=?,
    command_id=?,
    name=?,
    start_date=?,
    end_date=?,
    enable=?,
    ` + "`" + `repeat` + "`" + `=?
WHERE id = ?
`

type UpdateScheduleParams struct {
	TimeTypeID      int32     `json:"time_type_id"`
	IntervalDay     int32     `json:"interval_day"`
	IntervalSeconds int32     `json:"interval_seconds"`
	AtTime          string    `json:"at_time"`
	StartTime       string    `json:"start_time"`
	EndTime         string    `json:"end_time"`
	CommandID       int32     `json:"command_id"`
	Name            string    `json:"name"`
	StartDate       time.Time `json:"start_date"`
	EndDate         time.Time `json:"end_date"`
	Enable          bool      `json:"enable"`
	Repeat          bool      `json:"repeat"`
	ID              int32     `json:"id"`
}

func (q *Queries) UpdateSchedule(ctx context.Context, arg UpdateScheduleParams) error {
	_, err := q.exec(ctx, q.updateScheduleStmt, UpdateSchedule,
		arg.TimeTypeID,
		arg.IntervalDay,
		arg.IntervalSeconds,
		arg.AtTime,
		arg.StartTime,
		arg.EndTime,
		arg.CommandID,
		arg.Name,
		arg.StartDate,
		arg.EndDate,
		arg.Enable,
		arg.Repeat,
		arg.ID,
	)
	return err
}
