package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"log"
	"os"
	"path/filepath"
	"schedule/internal"
)

var (
	Trace *log.Logger
	Info  *log.Logger
	Error *log.Logger
)

func init() {
	newPath := filepath.Join(".", "log")
	err := os.MkdirAll(newPath, os.ModePerm)
	if err != nil {
		log.Fatal("can not create log folder")
	}
	file, err := os.OpenFile("./log/main.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("can not open log file")
	}
	Trace = log.New(os.Stdout, "TRACE: ", log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(io.MultiWriter(file, os.Stdout), "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(io.MultiWriter(file, os.Stdout), "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}
func main() {
	Info.Println("Starting...")
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/schedule/V1/all_schedules", internal.GetAllSchedulesRoute)
	r.POST("/schedule/V1/get_one_schedule", internal.GetScheduleOneRoute)
	r.POST("/schedule/V1/create_schedule", internal.CreateScheduleRoute)
	r.GET("/schedule/V1/all_commands", internal.GetAllCommandsRoute)
	r.POST("/schedule/V1/create_command", internal.CreateCommandRoute)
	r.POST("/schedule/V1/delete_command", internal.DeleteCommandRoute)
	r.POST("/schedule/V1/delete_schedule", internal.DeleteScheduleRoute)
	r.POST("/schedule/V1/update_schedule", internal.UpdateScheduleRoute)
	r.POST("/schedule/V1/update_command", internal.UpdateCommandRoute)
	r.GET("/schedule/V1/get_all_sub_schedules", internal.GetAllSubSchedulesRoute)

	r.Run(":9567")
}
