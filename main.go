package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Server struct {
	DB   *sql.DB
	GORM *gorm.DB
}

func ConnectToDatabase() (*gorm.DB, error) {
	dsn := "root:sujith10@tcp(localhost:3306)/test1?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}

	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {

	r := chi.NewRouter()
	// Route requests
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from home"))
	})
	r.Post("/alerts", WriteAlert)
	r.Get("/alerts/service_id={service_id}&start_ts={alert_ts}&end_ts={alert_end_ts}", ReadAlerts)
	// Server start
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", "8080"),
		Handler: r,
	}
	log.Println("Server started...")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(fmt.Sprintf("%+v", err))
	}

	db, err := ConnectToDatabase()
	if err != nil {
		panic("error" + err.Error())
	}
	fmt.Println(db.Name())
}

type Result struct {
	Alter_id string
	Err      error
}

type ReqData struct {
	AlertID     string `json:"alert_id"`
	ServiceID   string `json:"service_id"`
	ServiceName string `json:"service_name"`
	Model       string `json:"model"`
	AlertType   string `json:"alert_type"`
	AlertTS     string `json:"alert_ts"`
	Severity    string `json:"severity"`
	TeamSlack   string `json:"team_slack"`
}

type Data struct {
	ServiceID   string   `json:"service_id" gorm:"primaryKey"`
	ServiceName string   `json:"service_name"`
	Alerts      []Alerts `json:"alerts" gorm:"foreignKey:ServiceID;references:ServiceID"`
}

type Alerts struct {
	AlertID   string `json:"alert_id" gorm:"primaryKey"`
	Model     string `json:"model"`
	AlertType string `json:"alert_type"`
	AlertTs   string `json:"alert_ts"`
	Severity  string `json:"severity"`
	TeamSlack string `json:"team_slack"`
	ServiceID string `json:"service_id" gorm:"type:varchar(191);unique;index"`
}

// POST Request Handler (Write Alert)
func WriteAlert(w http.ResponseWriter, r *http.Request) {
	var data ReqData

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		fmt.Errorf(err.Error())
	}
	fmt.Println(data)

	db, err := ConnectToDatabase()
	if err != nil {
		panic("error" + err.Error())
	}
	fmt.Println(db.Name())

	err = db.AutoMigrate(&Data{})
	if err != nil {
		panic("Failed to auto-migrate the database: " + err.Error())
	}
	if err != nil {
		panic("Failed to auto-migrate the database: " + err.Error())
	}
	err = db.AutoMigrate(&Alerts{})
	if err != nil {
		panic("Failed to auto-migrate the database: " + err.Error())
	}

	datan := Data{
		ServiceID:   data.ServiceID,
		ServiceName: data.ServiceName,
	}
	alert1 := Alerts{
		AlertID:   data.AlertID,
		Model:     data.Model,
		AlertType: data.AlertType,
		AlertTs:   data.AlertTS,
		Severity:  data.Severity,
		TeamSlack: data.TeamSlack,
		ServiceID: data.ServiceID,
	}

	db.Create(datan)
	if err != nil {
		log.Fatal(err)
	}

	db.Create(alert1)
	if err != nil {
		log.Fatal(err)
	}
	result := Result{Alter_id: alert1.AlertID, Err: err}
	jsonResult, err := json.Marshal(result)

	w.WriteHeader(http.StatusOK)

	_, err = w.Write(jsonResult)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

// GET Request Handler (Read Alerts)
func ReadAlerts(w http.ResponseWriter, r *http.Request) {

	service_id := chi.URLParam(r, "service_id")
	alert_ts := chi.URLParam(r, "alert_ts")
	alert_end_ts := chi.URLParam(r, "alert_end_ts")
	fmt.Println(service_id, alert_end_ts, alert_ts)

	db, err := ConnectToDatabase()
	if err != nil {
		panic("error" + err.Error())
	}

	fmt.Println(db.Name())

	datan := Data{
		ServiceID: service_id,
	}
	db.Where("service_id = ?", service_id).Find(&datan)

	alert := []Alerts{}

	db.Where("service_id = ? BETWEEN ? AND ?", service_id, alert_ts, alert_end_ts).Find(&alert)
	datan.Alerts = alert
	fmt.Println(datan)
	w.WriteHeader(http.StatusOK)
	result := Result{Alter_id: alert[0].AlertID, Err: err}
	jsonResult, err := json.Marshal(result)
	_, err = w.Write(jsonResult)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
