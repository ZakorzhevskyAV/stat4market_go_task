package main

import (
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"net/http"
)

func SendEvent(w http.ResponseWriter, r *http.Request) {
	jsonmap := make(map[string]interface{})

	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("failed to read request body, err: %s", err)
		return
	}

	err = json.Unmarshal(data, &jsonmap)
	if err != nil {
		log.Printf("failed to unmarshal request body data into a map, err: %s", err)
		return
	}

	eventType := jsonmap["eventType"]
	userID := jsonmap["userID"]
	eventTime := jsonmap["eventTime"]
	payload := jsonmap["payload"]

	_, err = DBConn.Exec(
		`INSERT INTO TABLE events (eventID, eventType, userID, eventTime, payload) VALUES (?, ?, ?, ?, ?)`,
		int64(rand.Intn(10)), eventType, userID, eventTime, payload,
	)
	if err != nil {
		log.Printf("failed to insert rows in a table, err: %s", err)
		return
	}

	_, err = w.Write([]byte("data sent successfully"))
	if err != nil {
		log.Printf("failed to write a response, err: %s", err)
		return
	}
	log.Printf("data sent successfully")
}
