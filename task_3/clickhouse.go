package main

import (
	"database/sql"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func CHConnect() (*sql.DB, error) {
	conn := clickhouse.OpenDB(&clickhouse.Options{
		Addr: []string{fmt.Sprintf("%s:%d", os.Getenv("CLICKHOUSE_HOST"), 8123)},
		Auth: clickhouse.Auth{
			Database: os.Getenv("CLICKHOUSE_DB"),
			Username: os.Getenv("CLICKHOUSE_USER"),
			Password: os.Getenv("CLICKHOUSE_PASSWORD"),
		},
		Protocol: clickhouse.HTTP,
	})
	err := conn.Ping()
	if err != nil {
		return conn, err
	}
	log.Printf("connection successful")
	return conn, nil
}

func CreateTable(DBConn *sql.DB) error {
	_, err := DBConn.Exec(
		`CREATE TABLE events (
    eventID Int64,
    eventType String,
    userID Int64,
    eventTime DateTime,
    payload String
                    ) ENGINE = MergeTree
    ORDER BY (eventID, eventTime);`,
	)
	if err != nil {
		log.Printf("table failed to be created or already exists")
	}
	log.Printf("table creation successful")
	return nil
}

func FillTable(DBConn *sql.DB) error {
	var err error

	_, err = DBConn.Exec(
		`INSERT INTO TABLE events (eventID, eventType, userID, eventTime, payload) VALUES (?, ?, ?, ?, ?)`,
		1,
		"2",
		3,
		time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		"5",
	)
	if err != nil {
		log.Printf("failed to fill a table with a custom row")
		return err
	}
	log.Printf("custom test row inserted, sleeping for 1 second")
	time.Sleep(time.Second)

	for i := 0; i < 20; i++ {
		_, err = DBConn.Exec(
			`INSERT INTO TABLE events (eventID, eventType, userID, eventTime, payload) VALUES (?, ?, ?, ?, ?)`,
			int64(rand.Intn(10)),
			strconv.Itoa(rand.Intn(10)),
			int64(rand.Intn(100)),
			time.Unix(rand.Int63n(999999999)+1, 0),
			strconv.Itoa(rand.Intn(1000)),
		)
		if err != nil {
			log.Printf("failed to fill a table with row number %d, err: %s", i+1, err)
			return err
		}
		log.Printf("%s", time.Unix(rand.Int63n(999999999)+1, rand.Int63n(999999999)+1))
		log.Printf("row number %d inserted, sleeping for 1 second", i+1)
		time.Sleep(time.Second)
	}
	log.Printf("table filling successful")
	return nil
}
