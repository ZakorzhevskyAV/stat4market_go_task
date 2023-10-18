package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func GetRows(w http.ResponseWriter, r *http.Request) {
	eventType := r.URL.Query().Get("eventType")

	before := r.URL.Query().Get("before")
	beforeTime, err := time.Parse("2006-01-02", before)
	if err != nil {
		log.Printf("failed to parse a before time, err: %s", err)
		return
	}

	after := r.URL.Query().Get("after")
	afterTime, err := time.Parse("2006-01-02", after)
	if err != nil {
		log.Printf("failed to parse an after time, err: %s", err)
		return
	}

	DBRows, err := DBConn.Query(
		`SELECT * FROM events WHERE eventType = ? AND eventTime between ? AND ?`,
		eventType, beforeTime, afterTime,
	)
	if err != nil {
		log.Printf("failed to get rows from a table, err: %s", err)
		return
	}

	var row Row
	var rows []Row
	for DBRows.Next() {
		err = DBRows.Scan(&row.EventID, &row.EventType, &row.UserID, &row.EventTime, &row.Payload)
		if err != nil {
			log.Printf("failed to scan rows, err: %s", err)
			return
		}
		rows = append(rows, row)
	}

	log.Printf("rows: %s", rows)

	data, err := json.Marshal(rows)
	if err != nil {
		log.Printf("failed to marshal rows, err: %s", err)
		return
	}

	_, err = w.Write(data)
	if err != nil {
		log.Printf("failed to write a response, err: %s", err)
		return
	}
	log.Printf("data selected successfully")
}
