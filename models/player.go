package models

import (
	"database/sql"
	rl "github.com/gen2brain/raylib-go/raylib"
	"log"
	"sambragge/mymmo/util"
)

type Player struct {
	ID int64
	X int32
	Y int32
	speed int32
	height int32
	width int32
	color rl.Color
	*sql.DB
	online bool
}


func NewPlayer(db *sql.DB, x, y int32) (*Player, error) {
	p := &Player{
		X:x,
		Y:y,
		DB:db,
	}
	p.setConstants()

	res, err := p.Exec("INSERT INTO players(x, y, online) VALUES (?, ?, ?)", p.X, p.Y, true)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	p.ID = id

	log.Printf("-- WRITING %v TO FILE", id)

	if err := util.WriteToFile(id); err != nil {

		_, err2 := p.Exec("DELETE FROM players WHERE player_id = ?", id)
		util.HandleErr(err2, "player", "40")
		return nil, err
	}

	return p, nil
}

func ReturningPlayer(db *sql.DB, id int64) *Player {
	p := &Player{
		ID:id,
		DB:db,
	}
	p.setConstants()

	_, err := p.Exec("UPDATE players SET online = TRUE WHERE player_id = ?", p.ID)
	util.HandleErr(err, "plaeyer", "63")

	rows, err := p.Query("SELECT x, y FROM players WHERE player_id = ?", p.ID)
	util.HandleErr(err, "player", "60")
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&p.X, &p.Y)
	}

	return p
}

func(p *Player) setConstants(){
	p.height = 25
	p.width = 25
	p.speed = 10
	p.color = rl.Blue
}

func(p *Player) LogOff(){
	_, err := p.Exec("UPDATE players SET online = FALSE WHERE player_id = ?", p.ID)
	util.HandleErr(err, "player", "85")
}

func(p *Player) DrawAll(){
	rows, err := p.Query("SELECT x, y FROM players WHERE online = TRUE")
	util.HandleErr(err, "player", "56")
	defer rows.Close()
	for rows.Next() {
		var player Player
		err := rows.Scan(&player.X, &player.Y)
		util.HandleErr(err, "player", "61")
		player.setConstants()

		rl.DrawRectangle(player.X, player.Y, player.width, player.height, player.color)

	}
}

func(p *Player) Update(){

	changed := false

	if rl.IsKeyDown(rl.KeyA) {
		p.X -= p.speed
		changed = true
	}
	if rl.IsKeyDown(rl.KeyW) {
		p.Y -= p.speed
		changed = true
	}
	if rl.IsKeyDown(rl.KeyS) {
		p.Y += p.speed
		changed = true
	}
	if rl.IsKeyDown(rl.KeyD) {
		p.X += p.speed
		changed = true
	}

	if changed {
		_, err := p.Exec("UPDATE players SET x = ?, y = ? WHERE player_id = ?", p.X, p.Y, p.ID)
		util.HandleErr(err, "player", "97")
	}

}