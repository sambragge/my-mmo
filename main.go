package main

import (
	"database/sql"
	rl "github.com/gen2brain/raylib-go/raylib"
	_ "github.com/go-sql-driver/mysql"
	"sambragge/mymmo/client"
)



func main(){


	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/mymmo")
	if err != nil {
		panic(err.Error())
	}
	if err := db.Ping(); err != nil {
		panic(err.Error())
	}
	defer db.Close()

	c := client.Initialize(db)

	for !rl.WindowShouldClose() {

		c.Run()
	}

	c.HandleExit()
}

