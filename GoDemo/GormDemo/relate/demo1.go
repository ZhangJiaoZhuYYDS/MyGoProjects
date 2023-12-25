package main

import "GormDemo/initDB"

func main() {
	_db := initDB.DB
	println(_db)
}
