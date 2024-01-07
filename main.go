package main

func main() {
	listDrivers()
	db, err := openDatbase()
	if err == nil {
		db.Close()
	} else {
		panic(err)
	}
}
