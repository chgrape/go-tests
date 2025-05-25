package models

import "time"

type Todo struct {
	Id           int
	Title        string
	Completed    bool
	Created_at   time.Time
	Completed_at *time.Time
}
