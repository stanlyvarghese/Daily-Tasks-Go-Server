package main

import (
	"log"
)

//

//db, err := sql.Open("sqlite3", "app.db")

func main() {
	server, err := NewServer()

	if err != nil {
		log.Fatal(err)
	}

	testUserDatabase(server)

	server.Start()
}
