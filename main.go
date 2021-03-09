package main

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	cnf "main/config"
	c_auth "main/controller/auth"
	c_comm "main/controller/commons"
	c_mess "main/controller/messages"
	c_prof "main/controller/profile"
	c_sms "main/controller/sms"
	c_stud "main/controller/students"
	c_teach "main/controller/teachers"
	"net/http"
	"os"
	"runtime"
)

var (
	sep = string(os.PathSeparator)
)

type App struct {
	Router *mux.Router
	db     *sql.DB
	err    error
}

func init() {
	log.SetFlags(log.Lshortfile)
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	auth := c_auth.NewAuthController()
	teach := c_teach.NewTeachersController()
	stud := c_stud.NewStudentsController()
	comm := c_comm.NewCommonsController()
	sms := c_sms.NewSmsController()
	prof := c_prof.NewProfileController()
	mess := c_mess.NewMessagesController()
	router := mux.NewRouter()

	ath := router.PathPrefix("/auth").Subrouter()
	//auth.Use(authMiddleware)
	ath.HandleFunc("/login", auth.Login).Methods("POST")

	teachers := router.PathPrefix("/teachers").Subrouter()
	teachers.Use(authMiddleware)
	teachers.HandleFunc("/list", teach.GetList).Methods("POST")
	teachers.HandleFunc("/get", teach.GetDetail).Methods("POST")
	teachers.HandleFunc("/procs", teach.GetProcs).Methods("POST")

	students := router.PathPrefix("/students").Subrouter()
	students.Use(authMiddleware)
	students.HandleFunc("/list", stud.GetList).Methods("POST")
	students.HandleFunc("/get", stud.GetDetail).Methods("POST")
	students.HandleFunc("/procs", stud.GetProcs).Methods("POST")

	profile := router.PathPrefix("/profile").Subrouter()
	profile.Use(authMiddleware)
	profile.HandleFunc("/detail", prof.GetProfileDet).Methods("POST")
	profile.HandleFunc("/procs", prof.GetProcs).Methods("POST")

	common := router.PathPrefix("/com").Subrouter()
	common.Use(authMiddleware)
	common.HandleFunc("/dash", comm.GetDashboard).Methods("POST")
	common.HandleFunc("/changepass", comm.ChangePass).Methods("POST")

	messages := router.PathPrefix("/messages").Subrouter()
	messages.Use(authMiddleware)
	messages.HandleFunc("/list", mess.GetList).Methods("POST")
	messages.HandleFunc("/detail", mess.GetDetail).Methods("POST")
	messages.HandleFunc("/procs", mess.GetProcs).Methods("POST")
	messages.HandleFunc("/boxblink", mess.GetBoxBlink).Methods("POST")

	smses := router.PathPrefix("/sms").Subrouter()
	smses.Use(authMiddleware)
	smses.HandleFunc("/list", sms.GetList).Methods("POST")
	smses.HandleFunc("/get", sms.GetNums).Methods("POST")
	smses.HandleFunc("/send", sms.RouterIndex).Methods("POST")

	router.HandleFunc("/", HomeHandler)

	//crt := "server.crt"
	//key := "server.key"
	if cnf.DevMode == "http" {
		log.Fatal(http.ListenAndServe(":8202", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))
	} else {
		crt := "/etc/pki/tls/certs/domain.tld.cert"
		key := "/etc/pki/tls/private/domain.tld.key"
		log.Fatal(http.ListenAndServeTLS("127.0.0.1:8202", crt, key, handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("hello worlds")
}
