package internal

import (
	"github.com/gin-gonic/gin"
	"schedule/pkg/db"
)

// 檢查若有錯誤將錯誤信息返回給client
func ResponseErrorInfo(c *gin.Context, err error) {
	c.JSON(200, gin.H{
		"result": "fail",
		"error":  err.Error(),
	})
}

// request完成返回結果
func FinishReturnResult(c *gin.Context, result interface{}) {
	c.JSON(200, gin.H{
		"result": "ok",
		"data":   result,
	})
	return
}

func GetAllSchedulesRoute(c *gin.Context) {
	result, err := GetAllSchedules()
	if err != nil {
		ResponseErrorInfo(c, err)
		return
	}
	FinishReturnResult(c, result)
}

func GetScheduleOneRoute(c *gin.Context) {
	ID := RequestId{}
	err := c.Bind(&ID)
	if err != nil {
		ResponseErrorInfo(c, err)
		return
	}
	result, err := GetScheduleOne(ID.Id)
	Trace.Println(result)
	if err != nil {
		ResponseErrorInfo(c, err)
		return
	}
	FinishReturnResult(c, result)
}

func CreateScheduleRoute(c *gin.Context) {
	store := NewStore(MyDB)
	Schedule := RequestScheduleOne{}
	err := c.ShouldBindJSON(&Schedule)
	if err != nil {
		ResponseErrorInfo(c, err)
		return
	}
	result, err := store.CreateScheduleTX(Schedule.TimeTypeID, Schedule.IntervalDay, Schedule.IntervalSeconds, Schedule.AtTime,
		Schedule.StartTime, Schedule.EndTime, Schedule.CommandID, Schedule.Name, Schedule.StartDate, Schedule.EndDate,
		*Schedule.Enable, *Schedule.Repeat, Schedule.RepeatWeekday, Schedule.RepeatDay, Schedule.RepeatMonth)
	if err != nil {
		ResponseErrorInfo(c, err)
		return
	}
	FinishReturnResult(c, result)
}

func GetAllCommandsRoute(c *gin.Context) {
	result, err := GetCommands()
	if err != nil {
		ResponseErrorInfo(c, err)
		return
	}
	FinishReturnResult(c, result)

}

func CreateCommandRoute(c *gin.Context) {
	Command := db.Command{}
	err := c.Bind(&Command)
	if err != nil {
		ResponseErrorInfo(c, err)
		return
	}
	result, err := CreateCommand(Command.Command)
	if err != nil {
		ResponseErrorInfo(c, err)
		return
	}
	FinishReturnResult(c, result)

}

func DeleteCommandRoute(c *gin.Context) {
	RequestId := RequestId{}
	err := c.Bind(&RequestId)
	if err != nil {
		ResponseErrorInfo(c, err)
		return
	}
	err = DeleteCommand(RequestId.Id)
	if err != nil {
		ResponseErrorInfo(c, err)
		return
	}
	FinishReturnResult(c, "")
}

func DeleteScheduleRoute(c *gin.Context) {
	store := NewStore(MyDB)
	RequestId := RequestId{}
	err := c.Bind(&RequestId)
	if err != nil {
		ResponseErrorInfo(c, err)
		return
	}
	err = store.DeleteScheduleTx(RequestId.Id)
	if err != nil {
		ResponseErrorInfo(c, err)
		return
	}
	FinishReturnResult(c, "")
}

func UpdateScheduleRoute(c *gin.Context) {
	UpdateScheduleParam := RequestScheduleOne{}
	err := c.ShouldBindJSON(&UpdateScheduleParam)
	if err != nil {
		ResponseErrorInfo(c, err)
		return
	}
	err = UpdateScheduleRawTX(UpdateScheduleParam.ID, UpdateScheduleParam.TimeTypeID, UpdateScheduleParam.IntervalDay,
		UpdateScheduleParam.IntervalSeconds, UpdateScheduleParam.AtTime, UpdateScheduleParam.StartTime,
		UpdateScheduleParam.EndTime, UpdateScheduleParam.CommandID, UpdateScheduleParam.Name, UpdateScheduleParam.StartDate,
		UpdateScheduleParam.EndDate, *UpdateScheduleParam.Enable, *UpdateScheduleParam.Repeat, UpdateScheduleParam.RepeatWeekday,
		UpdateScheduleParam.RepeatDay, UpdateScheduleParam.RepeatMonth)
	if err != nil {
		ResponseErrorInfo(c, err)
		return
	}
	FinishReturnResult(c, "")
}

func UpdateCommandRoute(c *gin.Context) {
	UpdateCommandParam := db.UpdateCommandParams{}
	err := c.Bind(&UpdateCommandParam)
	if err != nil {
		ResponseErrorInfo(c, err)
		return
	}
	err = UpdateCommand(UpdateCommandParam.ID, UpdateCommandParam.Command)
	if err != nil {
		ResponseErrorInfo(c, err)
		return
	}
	FinishReturnResult(c, "")
}

func GetAllSubSchedulesRoute(c *gin.Context) {
	result, err := GetAllSubSchedules()
	if err != nil {
		ResponseErrorInfo(c, err)
		return
	}
	FinishReturnResult(c, result)

}
