package routes

import (
	"net/http"

	"github.com/AwatefMdev/graduation_project/controllers"
)

func CreateRoutes(mux *http.ServeMux, uc *controllers.UserController, jc *controllers.JobController) {
	mux.HandleFunc("/register", uc.Register)
	mux.HandleFunc("/login", uc.Login)

	mux.HandleFunc("/job", jc.Create)
	mux.HandleFunc("/job/", jc.Job)
	mux.HandleFunc("/job/feed", jc.Feed)
}
