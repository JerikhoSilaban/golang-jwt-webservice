package main

import (
	"DTSGolang/Kelas3/Sesi2Bagian2/database"
	"DTSGolang/Kelas3/Sesi2Bagian2/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8000")
}
