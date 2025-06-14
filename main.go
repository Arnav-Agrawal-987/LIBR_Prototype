package main

import (
	"net/http"

	"github.com/Arnav-Agrawal-987/Libr_Prototype/db"
	"github.com/Arnav-Agrawal-987/Libr_Prototype/router"
)

func main() {
	db.Init()
	r := router.Router()
	http.ListenAndServe(":8000", r)
}
