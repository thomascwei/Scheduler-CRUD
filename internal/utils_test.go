package internal

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/require"
	"sort"
	"testing"
	"time"
)

func TestGetAllSchedules(t *testing.T) {
	result, err := GetAllSchedules()
	require.NoError(t, err)
	require.NotEmpty(t, result)
	got := len(result)
	require.NotZero(t, got)
	rows, err := MyDB.Query("SELECT * FROM schedules")
	require.NoError(t, err)
	want := 0
	for rows.Next() {
		want++
	}
	require.Equal(t, want, got)
}

func TestGetScheduleOne(t *testing.T) {
	var MaxId int32
	row := MyDB.QueryRow("SELECT Max(id) FROM schedules ")
	err := row.Scan(&MaxId)
	require.NoError(t, err)
	result, err := GetScheduleOne(MaxId)
	require.NoError(t, err)
	require.NotEmpty(t, result)
	got := result.CommandID
	require.NotZero(t, got)
	var want int32
	row = MyDB.QueryRow("SELECT command_id FROM schedules where id = ?", MaxId)
	err = row.Scan(&want)
	require.NoError(t, err)
	require.Equal(t, want, got)
}

func TestCreateCommand(t *testing.T) {
	command := "test command"
	got, err := CreateCommand(command)
	require.NoError(t, err)
	require.NotZero(t, got)
	var want int64
	row := MyDB.QueryRow("SELECT MAX(id) FROM commands")
	err = row.Scan(&want)
	require.NoError(t, err)
	require.Equal(t, want, got)
	var WantString string
	row = MyDB.QueryRow("SELECT command FROM commands where id=?", got)
	err = row.Scan(&WantString)
	require.NoError(t, err)
	require.Equal(t, WantString, command)
}

func TestGetCommands(t *testing.T) {
	var want int
	row := MyDB.QueryRow("SELECT count(id) FROM commands")
	err := row.Scan(&want)
	require.NoError(t, err)
	commands, err := GetCommands()
	require.NoError(t, err)
	got := len(commands)
	require.Equal(t, want, got)

}

func TestDeleteCommand(t *testing.T) {
	var MaxId int32
	row := MyDB.QueryRow("SELECT MAX(id) FROM commands")
	err := row.Scan(&MaxId)
	require.NoError(t, err)
	err = DeleteCommand(MaxId)
	require.NoError(t, err)
	row = MyDB.QueryRow("SELECT Count(id) FROM commands where id=?", MaxId)
	var Zero int
	err = row.Scan(&Zero)
	require.Zero(t, Zero)
}

func TestDeleteScheduleTx(t *testing.T) {
	store := NewStore(MyDB)
	var MaxId int32
	row := MyDB.QueryRow("SELECT Max(id) FROM schedules ")
	err := row.Scan(&MaxId)
	require.NoError(t, err)
	err = store.DeleteScheduleTx(MaxId)
	require.NoError(t, err)
	// 成功刪除後應該找不到該帳號Id
	var got int
	row = MyDB.QueryRow("SELECT Count(id) FROM schedules where id=?", MaxId)
	err = row.Scan(&got)
	require.NoError(t, err)
	require.Zero(t, got)
	row = MyDB.QueryRow("SELECT Count(id) FROM repeat_day where schedule_id=?", MaxId)
	err = row.Scan(&got)
	require.NoError(t, err)
	require.Zero(t, got)
	row = MyDB.QueryRow("SELECT Count(id) FROM repeat_month where schedule_id=?", MaxId)
	err = row.Scan(&got)
	require.NoError(t, err)
	require.Zero(t, got)
	row = MyDB.QueryRow("SELECT Count(id) FROM repeat_weekday where schedule_id=?", MaxId)
	err = row.Scan(&got)
	require.NoError(t, err)
	require.Zero(t, got)
}

func TestCreateScheduleTX(t *testing.T) {
	store := NewStore(MyDB)
	t.Run("Create type2 schedule & create weekdays", func(t *testing.T) {
		StartDate := time.Now()
		EndDate := StartDate.Add(time.Hour * 24)
		RepeatWeekday := []string{"Tue", "Wed", "Sat"}
		sort.Strings(RepeatWeekday)
		got, err := store.CreateScheduleTX(2, -1, -1, "00:01:00", "00:03:00", "00:04:00", 1, "template1", StartDate, EndDate, true,
			false, RepeatWeekday, []string{}, []string{})
		require.NoError(t, err)
		require.NotZero(t, got)
		// 檢查當前最大的Id是否就是這次建立產生的
		var want int64
		row := MyDB.QueryRow("SELECT Max(id) FROM schedules ")
		err = row.Scan(&want)
		require.NoError(t, err)
		require.Equal(t, want, got)
		var ReturnWeekdays []string
		var SingleWeekday string

		rows, err := MyDB.Query("select weekday from repeat_weekday where schedule_id=?", got)
		require.NoError(t, err)

		for rows.Next() {
			err := rows.Scan(&SingleWeekday)
			require.NoError(t, err)
			ReturnWeekdays = append(ReturnWeekdays, SingleWeekday)
		}
		sort.Strings(ReturnWeekdays)
		require.Equal(t, RepeatWeekday, ReturnWeekdays)

		weekdays, err := GetWeekdaysByID(int32(got))
		require.NoError(t, err)
		sort.Strings(weekdays)
		require.Equal(t, RepeatWeekday, weekdays)
	})
	t.Run("Create type3 schedule & create days and create Months", func(t *testing.T) {
		StartDate := time.Now()
		EndDate := StartDate.Add(time.Hour * 24)
		RepeatDay := []string{"1", "11", "31"}
		RepeatMonth := []string{"Jan", "Feb", "Mar", "Apr", "May"}
		sort.Strings(RepeatDay)
		sort.Strings(RepeatMonth)

		got, err := store.CreateScheduleTX(3, -1, -1, "00:01:00", "00:03:00", "00:04:00", 1, "template1", StartDate, EndDate, true,
			false, []string{}, RepeatDay, RepeatMonth)
		require.NoError(t, err)
		require.NotEmpty(t, got)
		var want int64
		row := MyDB.QueryRow("SELECT Max(id) FROM schedules ")
		err = row.Scan(&want)
		require.NoError(t, err)
		require.Equal(t, want, got)

		var ReturnDays []string
		var SingleDay string
		rows, err := MyDB.Query("select day from repeat_day where schedule_id=?", got)
		require.NoError(t, err)
		for rows.Next() {
			err := rows.Scan(&SingleDay)
			require.NoError(t, err)
			ReturnDays = append(ReturnDays, SingleDay)
		}
		sort.Strings(ReturnDays)
		require.Equal(t, RepeatDay, ReturnDays)

		var ReturnMonths []string
		var SingleMonth string
		rows, err = MyDB.Query("select month from repeat_month where schedule_id=?", got)
		require.NoError(t, err)
		for rows.Next() {
			err := rows.Scan(&SingleMonth)
			require.NoError(t, err)
			ReturnMonths = append(ReturnMonths, SingleMonth)
		}
		sort.Strings(ReturnMonths)
		require.Equal(t, RepeatMonth, ReturnMonths)
	})
}

func TestUpdateCommand(t *testing.T) {
	var MaxId int32
	row := MyDB.QueryRow("SELECT MAX(id) FROM commands")
	err := row.Scan(&MaxId)
	require.NoError(t, err)
	Want := "Update From Test"
	err = UpdateCommand(MaxId, Want)
	require.NoError(t, err)
	var Got string
	row = MyDB.QueryRow("SELECT command FROM commands where id=? limit 1", MaxId)
	err = row.Scan(&Got)
	require.Equal(t, Want, Got)

}

func TestUpdateScheduleRawTX(t *testing.T) {
	var MaxId int32
	row := MyDB.QueryRow("SELECT MAX(id) FROM schedules")
	err := row.Scan(&MaxId)
	require.NoError(t, err)
	StartDate := time.Now()
	EndDate := StartDate.Add(time.Hour * 24)
	RepeatWeekday := []string{"Tue", "Wed", "DDDDD"}
	sort.Strings(RepeatWeekday)
	err = UpdateScheduleRawTX(MaxId, 2, -1, -1, "00:01:00", "00:03:00", "00:04:00", 1, "RawTX", StartDate, EndDate, true,
		false, RepeatWeekday, []string{}, []string{})
	require.NoError(t, err)

	var ReturnWeekdays []string
	var SingleWeekday string
	rows, err := MyDB.Query("select weekday from repeat_weekday where schedule_id=?", MaxId)
	require.NoError(t, err)
	for rows.Next() {
		err := rows.Scan(&SingleWeekday)
		require.NoError(t, err)
		ReturnWeekdays = append(ReturnWeekdays, SingleWeekday)
	}
	sort.Strings(ReturnWeekdays)
	require.Equal(t, RepeatWeekday, ReturnWeekdays)
}

func TestGetAllSubSchedules(t *testing.T) {
	result, err := GetAllSubSchedules()
	require.NoError(t, err)

	var want int
	row := MyDB.QueryRow("SELECT count(*) FROM repeat_weekday")
	err = row.Scan(&want)
	require.NoError(t, err)
	require.Equal(t, want, len(result.RepeatWeekday))
	row = MyDB.QueryRow("SELECT count(*) FROM repeat_month")
	err = row.Scan(&want)
	require.NoError(t, err)
	require.Equal(t, want, len(result.RepeatMonth))
	row = MyDB.QueryRow("SELECT count(*) FROM repeat_day")
	err = row.Scan(&want)
	require.NoError(t, err)
	require.Equal(t, want, len(result.RepeatDay))
}
