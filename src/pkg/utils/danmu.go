package utils

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/glebarez/go-sqlite"
)

func CreateDanmuFile(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}

	sqlStmt := `
	create table comment (
		id integer primary key autoincrement,
		player_name text,
		msg text,
		user_id integer,
		timestamp integer
	);
	create table info (
		start_time integer,
		end_time integer
	);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, err
	}

	stmt, err := db.Prepare("insert into info (start_time) values (?)")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(time.Now().Unix())
	if err != nil {
		return nil, err
	}

	return db, nil
}

func AddDanmuRecord(db *sql.DB, playerName string, msg string, userId int, timestamp int) error {
	stmt, err := db.Prepare("insert into comment(player_name, msg, user_id, timestamp) values (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(playerName, msg, userId, timestamp)
	if err != nil {
		return err
	}

	return nil
}

type ArtPlayerDanmu struct {
	Text string `json:"text"`
	Time int    `json:"time"`
}

func GetDanmuList(path string) ([]ArtPlayerDanmu, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer db.Close()

	var startTime int
	err = db.QueryRow("select start_time from info").Scan(&startTime)
	if err != nil && err == sql.ErrNoRows {
		fmt.Println(err.Error())
		return nil, err
	}

	rows, err := db.Query("select msg, timestamp from comment")
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer rows.Close()

	ret := make([]ArtPlayerDanmu, 0)
	for rows.Next() {
		var msg string
		var timestamp int
		err = rows.Scan(&msg, &timestamp)
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
		time := timestamp/1000 - startTime
		if time < 0 {
			time = 0
		}
		ret = append(ret, ArtPlayerDanmu{
			Text: msg,
			Time: time,
		})
	}
	return ret, nil
}
