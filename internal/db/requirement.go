package db

import "time"

type Requirement struct {
	Date   time.Time
	Client Client
	Name   any
}
