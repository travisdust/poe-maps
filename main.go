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

type Drop struct {
	MapId int `json:"mapId"`
	Count int `json:"count"`
}

type Run struct {
	Id int `db:"id",json:"id"`
	MapId int `db:"map_id",json:"mapId"`
	MapReturns []int `db:`
}

var db *sqlx.DB

func main() {

	db, err := sqlx.Connect("mysql", "root:root@(localhost:3306)/poe_maps")
	if err != nil {
		log.Fatalln(err)
	}


	r := gin.Default()
	r.GET("/maps", func(c *gin.Context) {

		maps := []Map{}
		err := db.Select(&maps, "SELECT id, name ,tier FROM map")
		if err != nil {
			log.Fatalln(err)
		}

		c.JSON(200, maps)
	})

	r.POST("/run", func(c *gin.Context) {
		var json struct {
			MapId int `json:"mapId"`
			Drops []Drop `json:"drops"`
		}

		c.BindJSON(&json)
		res , err := db.Exec("INSERT into run (map_id) VALUES (?)", json.MapId)
		if err != nil {
			log.Fatalln(err)
		}

		runId, _ := res.LastInsertId()
		for _, drop := range json.Drops {
			for i := 0; i < drop.Count; i++ {
				_, err := db.Exec("INSERT INTO `drop` (run_id, map_id) VALUES (?, ?)", runId, drop.MapId)
				if err != nil {
					log.Fatalln(err)
				}
			}
		}


		log.Print(json)
	})

	r.Run()
}