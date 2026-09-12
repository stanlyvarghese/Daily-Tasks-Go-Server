package main 

//

//db, err := sql.Open("sqlite3", "app.db")




func main() {

	server := NewServer()
	server.Start()
	
	
}
