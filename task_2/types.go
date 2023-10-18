package main

import "time"

type Row struct {
	EventID   int64     `json:"EventID"`
	EventType string    `json:"EventType"`
	UserID    int64     `json:"UserID"`
	EventTime time.Time `json:"EventTime"`
	Payload   string    `json:"Payload"`
}
