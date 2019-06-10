package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

type Map struct {
	Id int `db:"id",json:"id"`
	Name string `db:"name",json:"name"`
	Tier int `db:"tier",json:"tier"`
}
var db *sqlx.DB

func main() {

	db, err := sqlx.Connect("mysql", "root:root@(localhost:3306)/poe_maps")
	if err != nil {
		log.Fatalln(err)
	}


	r := gin.Default()
	r.GET("/", func(c *gin.Context) {

		maps := []Map{}
		err := db.Select(&maps, "SELECT id, name ,tier FROM map")
		if err != nil {
			log.Fatalln(err)
		}

		//mapsJson, _ := json.Marshal(maps)

		c.JSON(200, maps)
	})

	r.Run()
}