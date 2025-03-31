package entity

import "time"

type Book struct {
	ID        int64
	Title     string
	Author    string
	Published time.Time
}
