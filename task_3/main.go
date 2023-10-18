package main

import (
	"database/sql"
	"log"
	"net/http"
	"sync"
	"time"
)

var DBConn *sql.DB

func main() {
	var err error
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for {
			DBConn, err = CHConnect()
			if err != nil {
				log.Printf("failed to connect to ClickHouse, err: %s", err)
				time.Sleep(time.Second)
			} else {
				break
			}
		}
		wg.Done()
	}()
	wg.Wait()
	defer DBConn.Close()
	err = CreateTable(DBConn)
	if err != nil {
		log.Fatalf("failed to create a table, err: %s", err)
	}
	//err = FillTable(DBConn)
	//if err != nil {
	//	log.Fatalf("failed to fill a table, err: %s", err)
	//}
	http.HandleFunc("/", SendEvent)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
