package client

import (
	"database/sql"
	rl "github.com/gen2brain/raylib-go/raylib"
	"io/ioutil"
	"log"
	"os"
	"sambragge/mymmo/models"
	"sambragge/mymmo/util"
	"strconv"
	"strings"
)

const (
	SCREEN_WIDTH, SCREEN_HEIGHT int32 = 800, 600
	FPS int32 = 60
	TITLE string = "mymmo"
)

type Client struct {
	player *models.Player
}

func(c *Client) Run(){

	rl.BeginDrawing()

	rl.ClearBackground(rl.White)

	// update player
	go c.player.Update()

	// draw all players in database
	c.player.DrawAll()

	rl.EndDrawing()

}

func(c *Client) HandleExit(){

	c.player.LogOff()

	rl.CloseWindow()
}

func Initialize(db *sql.DB) *Client {
	c := &Client{}

	// check local data for user id
	file, err := os.Open(".data/player.txt")

	if err != nil {

		// general errors --
		// ...

		// file not found --
		if strings.Contains(err.Error(), "no such file or directory") {
			// make new player
			player, err := models.NewPlayer(db, SCREEN_WIDTH/2, SCREEN_HEIGHT/2)
			util.HandleErr(err, "client", "55")
			c.player = player

			log.Printf("-- SUCCESS CREATING NEW PLAYER WITH ID OF : %v", player.ID)
		}
	}else{
		// get player_id from local data
		bytes, err := ioutil.ReadAll(file)
		str := string(bytes)

		i, err := strconv.ParseInt(str, 0, 64)
		util.HandleErr(err, "client", "66")
		c.player = models.ReturningPlayer(db, i)

		log.Printf("-- SUCCESS RETRIEVING EXISTING USER WITH ID OF : %v", i)
	}
	defer file.Close()


	rl.InitWindow(SCREEN_WIDTH, SCREEN_HEIGHT, TITLE)
	rl.SetTargetFPS(FPS)

	return c
}