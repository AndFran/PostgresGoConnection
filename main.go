package main

func queryDatabase(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM person")
	if err == nil {
		for rows.Next() {
			var id int
			var firstName string
			var lastName string
			rows.Scan(&id, &firstName, &lastName)
			fmt.Println(id, firstName, lastName)
		}
	} else {
		log.Fatal(err.Error())
	}
}

func main() {
	listDrivers()
	db, err := openDatbase()
	if err == nil {
		queryDatabase(db)
		db.Close()
	} else {
		panic(err)
	}
}
