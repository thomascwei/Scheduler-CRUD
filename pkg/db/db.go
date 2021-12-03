// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.createCommandStmt, err = db.PrepareContext(ctx, CreateCommand); err != nil {
		return nil, fmt.Errorf("error preparing query CreateCommand: %w", err)
	}
	if q.createRepeatDaysStmt, err = db.PrepareContext(ctx, CreateRepeatDays); err != nil {
		return nil, fmt.Errorf("error preparing query CreateRepeatDays: %w", err)
	}
	if q.createRepeatMonthStmt, err = db.PrepareContext(ctx, CreateRepeatMonth); err != nil {
		return nil, fmt.Errorf("error preparing query CreateRepeatMonth: %w", err)
	}
	if q.createRepeatWeekdaysStmt, err = db.PrepareContext(ctx, CreateRepeatWeekdays); err != nil {
		return nil, fmt.Errorf("error preparing query CreateRepeatWeekdays: %w", err)
	}
	if q.createScheduleStmt, err = db.PrepareContext(ctx, CreateSchedule); err != nil {
		return nil, fmt.Errorf("error preparing query CreateSchedule: %w", err)
	}
	if q.deleteCommandStmt, err = db.PrepareContext(ctx, DeleteCommand); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteCommand: %w", err)
	}
	if q.deleteRepeatDayStmt, err = db.PrepareContext(ctx, DeleteRepeatDay); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteRepeatDay: %w", err)
	}
	if q.deleteRepeatMonthStmt, err = db.PrepareContext(ctx, DeleteRepeatMonth); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteRepeatMonth: %w", err)
	}
	if q.deleteRepeatWeekdayStmt, err = db.PrepareContext(ctx, DeleteRepeatWeekday); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteRepeatWeekday: %w", err)
	}
	if q.deleteScheduleStmt, err = db.PrepareContext(ctx, DeleteSchedule); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteSchedule: %w", err)
	}
	if q.getCommandsStmt, err = db.PrepareContext(ctx, GetCommands); err != nil {
		return nil, fmt.Errorf("error preparing query GetCommands: %w", err)
	}
	if q.getDaysStmt, err = db.PrepareContext(ctx, GetDays); err != nil {
		return nil, fmt.Errorf("error preparing query GetDays: %w", err)
	}
	if q.getMonthsStmt, err = db.PrepareContext(ctx, GetMonths); err != nil {
		return nil, fmt.Errorf("error preparing query GetMonths: %w", err)
	}
	if q.getScheduleStmt, err = db.PrepareContext(ctx, GetSchedule); err != nil {
		return nil, fmt.Errorf("error preparing query GetSchedule: %w", err)
	}
	if q.getWeekdaysStmt, err = db.PrepareContext(ctx, GetWeekdays); err != nil {
		return nil, fmt.Errorf("error preparing query GetWeekdays: %w", err)
	}
	if q.listRepeatDaysStmt, err = db.PrepareContext(ctx, ListRepeatDays); err != nil {
		return nil, fmt.Errorf("error preparing query ListRepeatDays: %w", err)
	}
	if q.listRepeatMonthsStmt, err = db.PrepareContext(ctx, ListRepeatMonths); err != nil {
		return nil, fmt.Errorf("error preparing query ListRepeatMonths: %w", err)
	}
	if q.listRepeatWeekdaysStmt, err = db.PrepareContext(ctx, ListRepeatWeekdays); err != nil {
		return nil, fmt.Errorf("error preparing query ListRepeatWeekdays: %w", err)
	}
	if q.listScheduleStmt, err = db.PrepareContext(ctx, ListSchedule); err != nil {
		return nil, fmt.Errorf("error preparing query ListSchedule: %w", err)
	}
	if q.updateCommandStmt, err = db.PrepareContext(ctx, UpdateCommand); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateCommand: %w", err)
	}
	if q.updateScheduleStmt, err = db.PrepareContext(ctx, UpdateSchedule); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateSchedule: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.createCommandStmt != nil {
		if cerr := q.createCommandStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createCommandStmt: %w", cerr)
		}
	}
	if q.createRepeatDaysStmt != nil {
		if cerr := q.createRepeatDaysStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createRepeatDaysStmt: %w", cerr)
		}
	}
	if q.createRepeatMonthStmt != nil {
		if cerr := q.createRepeatMonthStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createRepeatMonthStmt: %w", cerr)
		}
	}
	if q.createRepeatWeekdaysStmt != nil {
		if cerr := q.createRepeatWeekdaysStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createRepeatWeekdaysStmt: %w", cerr)
		}
	}
	if q.createScheduleStmt != nil {
		if cerr := q.createScheduleStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createScheduleStmt: %w", cerr)
		}
	}
	if q.deleteCommandStmt != nil {
		if cerr := q.deleteCommandStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteCommandStmt: %w", cerr)
		}
	}
	if q.deleteRepeatDayStmt != nil {
		if cerr := q.deleteRepeatDayStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteRepeatDayStmt: %w", cerr)
		}
	}
	if q.deleteRepeatMonthStmt != nil {
		if cerr := q.deleteRepeatMonthStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteRepeatMonthStmt: %w", cerr)
		}
	}
	if q.deleteRepeatWeekdayStmt != nil {
		if cerr := q.deleteRepeatWeekdayStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteRepeatWeekdayStmt: %w", cerr)
		}
	}
	if q.deleteScheduleStmt != nil {
		if cerr := q.deleteScheduleStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteScheduleStmt: %w", cerr)
		}
	}
	if q.getCommandsStmt != nil {
		if cerr := q.getCommandsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getCommandsStmt: %w", cerr)
		}
	}
	if q.getDaysStmt != nil {
		if cerr := q.getDaysStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getDaysStmt: %w", cerr)
		}
	}
	if q.getMonthsStmt != nil {
		if cerr := q.getMonthsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getMonthsStmt: %w", cerr)
		}
	}
	if q.getScheduleStmt != nil {
		if cerr := q.getScheduleStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getScheduleStmt: %w", cerr)
		}
	}
	if q.getWeekdaysStmt != nil {
		if cerr := q.getWeekdaysStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getWeekdaysStmt: %w", cerr)
		}
	}
	if q.listRepeatDaysStmt != nil {
		if cerr := q.listRepeatDaysStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listRepeatDaysStmt: %w", cerr)
		}
	}
	if q.listRepeatMonthsStmt != nil {
		if cerr := q.listRepeatMonthsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listRepeatMonthsStmt: %w", cerr)
		}
	}
	if q.listRepeatWeekdaysStmt != nil {
		if cerr := q.listRepeatWeekdaysStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listRepeatWeekdaysStmt: %w", cerr)
		}
	}
	if q.listScheduleStmt != nil {
		if cerr := q.listScheduleStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listScheduleStmt: %w", cerr)
		}
	}
	if q.updateCommandStmt != nil {
		if cerr := q.updateCommandStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateCommandStmt: %w", cerr)
		}
	}
	if q.updateScheduleStmt != nil {
		if cerr := q.updateScheduleStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateScheduleStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                       DBTX
	tx                       *sql.Tx
	createCommandStmt        *sql.Stmt
	createRepeatDaysStmt     *sql.Stmt
	createRepeatMonthStmt    *sql.Stmt
	createRepeatWeekdaysStmt *sql.Stmt
	createScheduleStmt       *sql.Stmt
	deleteCommandStmt        *sql.Stmt
	deleteRepeatDayStmt      *sql.Stmt
	deleteRepeatMonthStmt    *sql.Stmt
	deleteRepeatWeekdayStmt  *sql.Stmt
	deleteScheduleStmt       *sql.Stmt
	getCommandsStmt          *sql.Stmt
	getDaysStmt              *sql.Stmt
	getMonthsStmt            *sql.Stmt
	getScheduleStmt          *sql.Stmt
	getWeekdaysStmt          *sql.Stmt
	listRepeatDaysStmt       *sql.Stmt
	listRepeatMonthsStmt     *sql.Stmt
	listRepeatWeekdaysStmt   *sql.Stmt
	listScheduleStmt         *sql.Stmt
	updateCommandStmt        *sql.Stmt
	updateScheduleStmt       *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                       tx,
		tx:                       tx,
		createCommandStmt:        q.createCommandStmt,
		createRepeatDaysStmt:     q.createRepeatDaysStmt,
		createRepeatMonthStmt:    q.createRepeatMonthStmt,
		createRepeatWeekdaysStmt: q.createRepeatWeekdaysStmt,
		createScheduleStmt:       q.createScheduleStmt,
		deleteCommandStmt:        q.deleteCommandStmt,
		deleteRepeatDayStmt:      q.deleteRepeatDayStmt,
		deleteRepeatMonthStmt:    q.deleteRepeatMonthStmt,
		deleteRepeatWeekdayStmt:  q.deleteRepeatWeekdayStmt,
		deleteScheduleStmt:       q.deleteScheduleStmt,
		getCommandsStmt:          q.getCommandsStmt,
		getDaysStmt:              q.getDaysStmt,
		getMonthsStmt:            q.getMonthsStmt,
		getScheduleStmt:          q.getScheduleStmt,
		getWeekdaysStmt:          q.getWeekdaysStmt,
		listRepeatDaysStmt:       q.listRepeatDaysStmt,
		listRepeatMonthsStmt:     q.listRepeatMonthsStmt,
		listRepeatWeekdaysStmt:   q.listRepeatWeekdaysStmt,
		listScheduleStmt:         q.listScheduleStmt,
		updateCommandStmt:        q.updateCommandStmt,
		updateScheduleStmt:       q.updateScheduleStmt,
	}
}