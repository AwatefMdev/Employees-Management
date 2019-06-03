package main

import (
	"log"
	"net/http"
	"os"

	"github.com/AwatefMdev/graduation_project/controllers"
	"github.com/AwatefMdev/graduation_project/routes"
	"github.com/AwatefMdev/graduation_project/utils/caching"
	"github.com/AwatefMdev/graduation_project/utils/database"
)

func main() {
	db, err := database.Connect(os.Getenv("PGUSER"), os.Getenv("PGPASS"), os.Getenv("PGDB"), os.Getenv("PGHOST"), os.Getenv("PGPORT"))
	if err != nil {
		log.Fatal(err)
	}
	cache := &caching.Redis{
		Client: caching.Connect(os.Getenv("REDIS_ADDR"), os.Getenv("REDIS_PASSWORD"), 0),
	}

	employeeController := controllers.NewEmployeeController(db, cache)
	attendanceController := controllers.NewAttendanceController(db, cache)
	toolsController := controllers.NewToolsController(db, cache)
	leavesController := controllers.NewLeavesController(db, cache)
	rolesController := controllers.NewRolesController(db, cache)
	meeting_roomController := controllers.NewMeetingRoomController(db, cache)
	trainingController := controllers.NewTrainingController(db, cache)
	employee_meeting_roomController := controllers.NewEmployeeMeetingRoomController(db, cache)
	employee_trainingController := controllers.NewEmloyeeTrainingController(db, cache)
	parkingController := controllers.NewParkingController(db, cache)

	mux := http.NewServeMux()
	routes.CreateRoutes(mux, employeeController,attendanceController ,toolsController ,leavesController,rolesController ,meeting_roomController ,trainingController,employee_meeting_roomController ,employee_trainingController ,parkingController)

	if err := http.ListenAndServe(":8000", mux); err != nil {
		log.Fatal(err)
	}
}
