package internal

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/spf13/viper"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"schedule/pkg/db"
	"time"
)

var (
	file, _  = os.OpenFile("./log/utils.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	Trace    = log.New(os.Stdout, "TRACE: ", log.Ldate|log.Ltime|log.Lshortfile)
	Info     = log.New(io.MultiWriter(file, os.Stdout), "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error    = log.New(io.MultiWriter(file, os.Stdout), "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	txoption = sql.TxOptions{
		Isolation: 6,
		ReadOnly:  false,
	}

	DBconfig     = LoadConfig("./config")
	DBConnection = fmt.Sprintf("%s:%s@tcp(%s:%s)/schedule?charset=utf8&parseTime=true&parseTime=true&loc=Local",
		DBconfig.User, DBconfig.Password, DBconfig.Host, DBconfig.Port)
	//DBConnection = "thomas:123456@tcp(host.docker.internal:3306)/schedule?charset=utf8&parseTime=true"
	MyDB, _ = sql.Open("mysql", DBConnection)
	ctx     = context.Background()
	queries = db.New(MyDB)
)

// 讀專案中的config檔
func LoadConfig(mypath string) (config Config) {
	// 若有同名環境變量則使用環境變量
	viper.AutomaticEnv()
	viper.AddConfigPath(mypath)
	// 為了讓執行test也能讀到config添加索引路徑
	wd, err := os.Getwd()
	parent := filepath.Dir(wd)
	viper.AddConfigPath(path.Join(parent, mypath))
	viper.SetConfigName("db")
	viper.SetConfigType("yaml")
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal("can not load config: " + err.Error())
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatal("can not load config: " + err.Error())
	}
	Trace.Printf("%+v\n", config)
	return
}

func GetAllSchedules() ([]db.Schedule, error) {
	result, err := queries.ListSchedule(ctx)
	if err != nil {
		Error.Println(err.Error())
		return nil, err
	}
	return result, nil
}

func GetScheduleOne(id int32) (ResponseScheduleOne, error) {
	schedule, err := queries.GetSchedule(ctx, id)
	if err != nil {
		Error.Println(err.Error())
		return ResponseScheduleOne{}, err
	}
	// TimeTypeID:0 -> one shot
	// TimeTypeID:1 -> 指定日期範圍內的指定時間區間依時間間隔重複, 不用額外存
	// TimeTypeID:2 -> 指定星期幾範圍內指定時間區間依時間間隔重複, 要另外存星期幾
	// TimeTypeID:3 -> 指定月份範圍內指定日期範圍內指定時間區間依時間間隔重複, 要另外存月與日
	result := ResponseScheduleOne{}
	result.Schedule = schedule
	switch schedule.TimeTypeID {
	case 2:
		rows, err := GetWeekdaysByID(id)
		if err != nil {
			Error.Println(err.Error())
			return ResponseScheduleOne{}, err
		}
		result.RepeatWeekday = rows
	case 3:
		rows1, err := GetMonthsByID(id)
		if err != nil {
			Error.Println(err.Error())
			return ResponseScheduleOne{}, err
		}
		result.RepeatMonth = rows1
		rows2, err := GetDaysByID(id)
		if err != nil {
			Error.Println(err.Error())
			return ResponseScheduleOne{}, err
		}
		result.RepeatDay = rows2
	}
	return result, nil
}

func GetWeekdaysByID(id int32) ([]string, error) {
	var result []string
	rows, err := queries.GetWeekdays(ctx, id)
	if err != nil {
		Error.Println(err)
		return nil, err
	}
	for _, row := range rows {
		result = append(result, row.Weekday)
	}
	return result, nil
}

func GetDaysByID(id int32) ([]string, error) {
	var result []string
	rows, err := queries.GetDays(ctx, id)
	if err != nil {
		Error.Println(err)
		return nil, err
	}
	for _, row := range rows {
		result = append(result, row.Day)
	}
	return result, nil
}

func GetMonthsByID(id int32) ([]string, error) {
	var result []string
	rows, err := queries.GetMonths(ctx, id)
	if err != nil {
		Error.Println(err)
		return nil, err
	}
	for _, row := range rows {
		result = append(result, row.Month)
	}
	return result, nil
}

func GetCommands() ([]db.Command, error) {
	commands, err := queries.GetCommands(ctx)
	if err != nil {
		Error.Println(err)
		return nil, err
	}
	return commands, nil
}

func CreateCommand(command string) (int64, error) {
	result, err := queries.CreateCommand(ctx, command)
	if err != nil {
		Error.Println(err)
		return -1, err
	}
	LID, err := result.LastInsertId()
	if err != nil {
		Error.Println(err)
		return -1, err
	}
	return LID, nil
}

func DeleteCommand(id int32) error {
	err := queries.DeleteCommand(ctx, id)
	if err != nil {
		Error.Println(err)
		return err
	}
	return nil
}

// Store defines all functions to execute db queries and transactions
type Store struct {
	*db.Queries
	TransDB *sql.DB
}

// NewStore creates a new store
func NewStore(sdb *sql.DB) *Store {
	return &Store{
		Queries: db.New(sdb),
		TransDB: sdb,
	}
}

// ExecTx executes a function within a database transaction
func (s *Store) execTx(ctx context.Context, fn func(*db.Queries) error) error {
	tx, err := s.TransDB.BeginTx(ctx, &txoption)
	if err != nil {
		return err
	}
	q := db.New(tx)
	err = fn(q)
	if err != nil {
		rbErr := tx.Rollback()
		if rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)

		}
		return err
	}
	return tx.Commit()
}

// DeleteScheduleTx delete all related data within a database transaction
func (s *Store) DeleteScheduleTx(id int32) error {

	err := s.execTx(ctx, func(q *db.Queries) error {
		// 連續性操作實作於此
		err := q.DeleteRepeatDay(ctx, id)
		if err != nil {
			return err
		}
		err = q.DeleteRepeatMonth(ctx, id)
		if err != nil {
			return err
		}

		err = q.DeleteRepeatWeekday(ctx, id)
		if err != nil {
			return err
		}

		err = q.DeleteSchedule(ctx, id)
		if err != nil {
			return err
		}

		return nil
	})

	return err
}

// CreateScheduleTx all related data within a database transaction
func (s *Store) CreateScheduleTX(TimeTypeId int32, IntervalDay int32, IntervalSeconds int32, AtTime string, StartTime string,
	EndTime string, CommandId int32, Name string, StartDate time.Time, EndDate time.Time, Enable bool, Repeat bool,
	RepeatWeekday []string, RepeatDay []string, RepeatMonth []string) (int64, error) {
	var LID int64
	schedule := db.CreateScheduleParams{
		TimeTypeID:      TimeTypeId,
		IntervalDay:     IntervalDay,
		IntervalSeconds: IntervalSeconds,
		AtTime:          AtTime,
		StartTime:       StartTime,
		EndTime:         EndTime,
		CommandID:       CommandId,
		Name:            Name,
		StartDate:       StartDate,
		EndDate:         EndDate,
		Enable:          Enable,
		Repeat:          Repeat,
	}
	err := s.execTx(ctx, func(q *db.Queries) error {
		// 連續性操作實作於此
		result, err := queries.CreateSchedule(ctx, schedule)
		if err != nil {
			Error.Println(err)
			return err
		}
		LID, err = result.LastInsertId()
		if err != nil {
			Error.Println(err)
			return err
		}
		if schedule.TimeTypeID == 2 {
			for _, weekday := range RepeatWeekday {
				Params := db.CreateRepeatWeekdaysParams{
					ScheduleID: int32(LID),
					Weekday:    weekday,
				}
				err = queries.CreateRepeatWeekdays(ctx, Params)
				if err != nil {
					Error.Println(err)
					return err
				}
			}
		} else if schedule.TimeTypeID == 3 {
			for _, day := range RepeatDay {
				Params := db.CreateRepeatDaysParams{
					ScheduleID: int32(LID),
					Day:        day,
				}
				err = queries.CreateRepeatDays(ctx, Params)
				if err != nil {
					Error.Println(err)
					return err
				}
			}
			for _, month := range RepeatMonth {
				Params := db.CreateRepeatMonthParams{
					ScheduleID: int32(LID),
					Month:      month,
				}
				err = queries.CreateRepeatMonth(ctx, Params)
				if err != nil {
					Error.Println(err)
					return err
				}
			}
		}
		return err
	})
	return LID, err
}

// UpdateScheduleTx all related data within a database transaction
// Deprecated: generate locking.
func (s *Store) UpdateScheduleTX(Id int32, TimeTypeId int32, IntervalDay int32, IntervalSeconds int32, AtTime string, StartTime string,
	EndTime string, CommandId int32, Name string, StartDate time.Time, EndDate time.Time, Enable bool, Repeat bool,
	RepeatWeekday []string, RepeatDay []string, RepeatMonth []string) error {
	schedule := db.UpdateScheduleParams{
		ID:              Id,
		TimeTypeID:      TimeTypeId,
		IntervalDay:     IntervalDay,
		IntervalSeconds: IntervalSeconds,
		AtTime:          AtTime,
		StartTime:       StartTime,
		EndTime:         EndTime,
		CommandID:       CommandId,
		Name:            Name,
		StartDate:       StartDate,
		EndDate:         EndDate,
		Enable:          Enable,
		Repeat:          Repeat,
	}
	err := s.execTx(ctx, func(q *db.Queries) error {
		// 連續性操作實作於此
		// 此ID其他表內容先全刪
		err := q.DeleteRepeatWeekday(ctx, schedule.ID)
		if err != nil {
			Error.Println(err)
			return err
		}
		err = q.DeleteRepeatMonth(ctx, schedule.ID)
		if err != nil {
			Error.Println(err)
			return err
		}
		err = q.DeleteRepeatDay(ctx, schedule.ID)
		if err != nil {
			Error.Println(err)
			return err
		}
		// 修改schedule
		err = q.UpdateSchedule(ctx, schedule)
		if err != nil {
			Error.Println(err)
			return err
		}
		Trace.Println("finish UpdateSchedule")
		// 重新建其他關連表
		if schedule.TimeTypeID == 2 {
			for _, weekday := range RepeatWeekday {
				Params := db.CreateRepeatWeekdaysParams{
					ScheduleID: schedule.ID,
					Weekday:    weekday,
				}
				Trace.Println("準備新增weekday")
				Trace.Println(Params)
				err = queries.CreateRepeatWeekdays(ctx, Params)
				if err != nil {
					Error.Println(err)
					return err
				}
				Trace.Println("完成新增weekday")
			}
		} else if schedule.TimeTypeID == 3 {
			for _, day := range RepeatDay {
				Params := db.CreateRepeatDaysParams{
					ScheduleID: schedule.ID,
					Day:        day,
				}
				err = queries.CreateRepeatDays(ctx, Params)
				if err != nil {
					Error.Println(err)
					return err
				}
			}
			for _, month := range RepeatMonth {
				Params := db.CreateRepeatMonthParams{
					ScheduleID: schedule.ID,
					Month:      month,
				}
				err = queries.CreateRepeatMonth(ctx, Params)
				if err != nil {
					Error.Println(err)
					return err
				}
			}
		}
		return err
	})
	return err
}

// using database/sql BeginTx function to implement transaction
func UpdateScheduleRawTX(Id int32, TimeTypeId int32, IntervalDay int32, IntervalSeconds int32, AtTime string, StartTime string,
	EndTime string, CommandId int32, Name string, StartDate time.Time, EndDate time.Time, Enable bool, Repeat bool,
	RepeatWeekday []string, RepeatDay []string, RepeatMonth []string) error {
	schedule := db.UpdateScheduleParams{
		ID:              Id,
		TimeTypeID:      TimeTypeId,
		IntervalDay:     IntervalDay,
		IntervalSeconds: IntervalSeconds,
		AtTime:          AtTime,
		StartTime:       StartTime,
		EndTime:         EndTime,
		CommandID:       CommandId,
		Name:            Name,
		StartDate:       StartDate,
		EndDate:         EndDate,
		Enable:          Enable,
		Repeat:          Repeat,
	}
	tx, err := MyDB.BeginTx(ctx, &txoption)
	if err != nil {
		Error.Println(err)
		tx.Rollback()
		return err
	}
	// 連續性操作實作於此
	// 此ID其他表內容先全刪
	_, err = tx.Exec(db.DeleteRepeatWeekday, schedule.ID)
	if err != nil {
		Error.Println(err)
		tx.Rollback()
		return err
	}
	_, err = tx.Exec(db.DeleteRepeatMonth, schedule.ID)
	if err != nil {
		Error.Println(err)
		tx.Rollback()
		return err
	}
	_, err = tx.Exec(db.DeleteRepeatDay, schedule.ID)
	if err != nil {
		Error.Println(err)
		tx.Rollback()
		return err
	}
	// 修改schedule
	_, err = tx.Exec(db.UpdateSchedule, schedule.TimeTypeID, schedule.IntervalDay, schedule.IntervalSeconds,
		schedule.AtTime, schedule.StartTime, schedule.EndDate, schedule.CommandID, schedule.Name,
		schedule.StartDate, schedule.EndDate, schedule.Enable, schedule.Repeat, schedule.ID)
	if err != nil {
		Error.Println(err)
		tx.Rollback()
		return err
	}
	// 重新建其他關連表
	if schedule.TimeTypeID == 2 {
		for _, weekday := range RepeatWeekday {
			_, err = tx.Exec(db.CreateRepeatWeekdays, schedule.ID, weekday)
			if err != nil {
				Error.Println(err)
				tx.Rollback()
				return err
			}
		}
	} else if schedule.TimeTypeID == 3 {
		for _, day := range RepeatDay {
			_, err = tx.Exec(db.CreateRepeatDays, schedule.ID, day)
			if err != nil {
				Error.Println(err)
				tx.Rollback()
				return err
			}
		}
		for _, month := range RepeatMonth {
			_, err = tx.Exec(db.CreateRepeatMonth, schedule.ID, month)
			if err != nil {
				Error.Println(err)
				tx.Rollback()
				return err
			}
		}
	}
	return tx.Commit()
}

func UpdateCommand(Id int32, Command string) error {
	err := queries.UpdateCommand(ctx, db.UpdateCommandParams{
		Command: Command,
		ID:      Id,
	})
	if err != nil {
		Error.Println(err)
		return err
	}
	return nil
}

func GetAllSubSchedules() (SubSchedules, error) {
	ResultSubSchedules := SubSchedules{}
	RepeatWeekdays, err := queries.ListRepeatWeekdays(ctx)
	if err != nil {
		Error.Println(err)
		return SubSchedules{}, err
	}
	ResultSubSchedules.RepeatWeekday = RepeatWeekdays

	RepeatMonth, err := queries.ListRepeatMonths(ctx)
	if err != nil {
		Error.Println(err)
		return SubSchedules{}, err
	}
	ResultSubSchedules.RepeatMonth = RepeatMonth

	RepeatDay, err := queries.ListRepeatDays(ctx)
	if err != nil {
		Error.Println(err)
		return SubSchedules{}, err
	}
	ResultSubSchedules.RepeatDay = RepeatDay

	return ResultSubSchedules, err
}
