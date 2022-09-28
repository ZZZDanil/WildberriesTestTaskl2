package server

import (
	"business"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

var eventDB business.EventDB

func Run() {

	eventDB = *business.CreteEventDB()

	fmt.Println("ssss")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("+++")
		fmt.Println(r.URL.Path)
		fmt.Println(r.URL.Query())
		switch r.Method {
		case "POST":
			{
				switch r.URL.Path {
				case "/create_event":
					{
						POST_create_event(w, r)
					}
				case "/update_event":
					{
						POST_update_event(w, r)
					}
				case "/delete_event":
					{
						POST_delete_event(w, r)
					}
				}
			}
		case "GET":
			{
				switch r.URL.Path {
				case "/events_for_day":
					{
						GET_events_for_day(w, r)
					}
				case "/events_for_week":
					{
						GET_events_for_week(w, r)
					}
				case "/events_for_month":
					{
						GET_events_for_month(w, r)
					}
				}
			}
		}
	})
	log.Println("Запуск веб-сервера на http://127.0.0.1:4000")
	err := http.ListenAndServe(":4000", nil)
	log.Fatal(err)
}

func POST_create_event(w http.ResponseWriter, r *http.Request) {
	data := business.AddData{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err == nil {
		eventDB.Add(data.User_id, data.Event_name, data.Date)
	}
}
func POST_update_event(w http.ResponseWriter, r *http.Request) {

	data := business.UpdateData{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err == nil {
		eventDB.Update(data.User_id, data.Event_name, data.Date)
	}

}
func POST_delete_event(w http.ResponseWriter, r *http.Request) {
	data := business.DeleteData{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err == nil {
		eventDB.Delete(data.User_id, data.Event_name, data.Date)
	}
}
func GET_events_for_day(w http.ResponseWriter, r *http.Request) {
	m := r.URL.Query()
	user_id, ok1 := m["user_id"]
	event_name, ok2 := m["event_name"]

	if ok1 && ok2 {
		t := time.Now()
		eventDB.GetByRange(
			user_id[0], event_name[0],
			t.Truncate(time.Duration(t.Local().Day())).Format(business.Format),
			t.Format(business.Format))
	}
}
func GET_events_for_week(w http.ResponseWriter, r *http.Request) {
	m := r.URL.Query()
	user_id, ok1 := m["user_id"]
	event_name, ok2 := m["event_name"]

	if ok1 && ok2 {
		t := time.Now()
		eventDB.GetByRange(
			user_id[0], event_name[0],
			t.Truncate(time.Duration(t.Local().Weekday())).Format(business.Format),
			t.Format(business.Format))
	}
}
func GET_events_for_month(w http.ResponseWriter, r *http.Request) {
	m := r.URL.Query()
	user_id, ok1 := m["user_id"]
	event_name, ok2 := m["event_name"]

	if ok1 && ok2 {
		t := time.Now()
		eventDB.GetByRange(
			user_id[0], event_name[0],
			t.Truncate(time.Duration(t.Local().Month())).Format(business.Format),
			t.Format(business.Format))
	}
}
