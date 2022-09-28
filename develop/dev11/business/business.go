package business

import (
	"fmt"
	"time"
)

const Format = "2006-01-02T15:04:05.000Z"

type Event struct {
	id, user_id, event_name string
	date                    time.Time
}

func generateId() string {
	return fmt.Sprint(time.Now().Unix())
}
func convertTime(date string) (time.Time, error) {
	return time.Parse(Format, date)
}

func (e *Event) ChangeId() {
	e.id = generateId()
}

func CreateEvent(user_id, event_name, date string) (*Event, error) {

	t, err := convertTime(date)
	if err != nil {
		return nil, err
	}
	return &Event{
		id:         generateId(),
		user_id:    user_id,
		event_name: event_name,
		date:       t,
	}, nil
}

type EventDB struct {
	Events map[string]Event
}

func CreteEventDB() *EventDB {
	e := &EventDB{}
	e.Events = make(map[string]Event)
	return e
}

type AddData struct {
	User_id, Event_name, Date string
}

func (e *EventDB) Add(user_id, event_name, date string) {
	ev, err := CreateEvent(user_id, event_name, date)
	if err == nil {
		for _, ok := e.Events[ev.id]; !ok; {
			ev.ChangeId()
			_, ok = e.Events[ev.id]
		}
		e.Events[ev.id] = *ev
	}
}
func (e *EventDB) Get(user_id, event_name string) *Event {
	for _, ev := range e.Events {
		if ev.user_id == user_id && ev.event_name == event_name {
			return &ev
		}
	}
	return nil
}
func (e *EventDB) GetByRange(user_id, event_name, from, to string) *Event {
	f, err1 := convertTime(from)
	t, err2 := convertTime(to)
	if err1 == nil && err2 == nil {
		for _, ev := range e.Events {
			if ev.user_id == user_id && ev.event_name == event_name && ev.date.After(f) && ev.date.Before(t) {
				return &ev
			}
		}
	}
	return nil
}

type UpdateData struct {
	User_id, Event_name, Date string
}

func (e *EventDB) Update(user_id, event_name, date string) {
	t, err := convertTime(date)
	if err == nil {
		for _, ev := range e.Events {
			if ev.user_id == user_id && ev.event_name == event_name {
				ev.date = t
			}
		}
	}
}

type DeleteData struct {
	User_id, Event_name, Date string
}

func (e *EventDB) Delete(user_id, event_name, date string) {
	t, err := convertTime(date)
	if err == nil {
		for _, ev := range e.Events {
			if ev.user_id == user_id && ev.event_name == event_name && ev.date == t {
				delete(e.Events, ev.id)
			}
		}
	}
}
