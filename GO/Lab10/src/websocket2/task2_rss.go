package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mmcdole/gofeed"
)

const (
	password string = "Je2dTYr6"
	login    string = "iu9networkslabs"
	host     string = "students.yss.su"
	dbname   string = "iu9networkslabs"
)

func main() {
	db, err := sql.Open("mysql", login+":"+password+"@tcp("+host+")/"+dbname)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	for {
		add := ""
		fmt.Scan(&add)
		switch {
		case add == "update":
			fp := gofeed.NewParser()
			feed, _ := fp.ParseURL("https://vesti-k.ru/rss/")
			for _, item := range feed.Items {

				_, err1 := db.Exec("INSERT INTO iu9networkslabs.iu9Talgat (title,description,date,author,time) VALUES  (?, ?, ?, ?, ?)",
					item.Title, item.Description, item.PublishedParsed, "-", time.Now())
				if err1 != nil {
					panic(err1)
				}
			}
		case add == "clear":
			db.Query("TRUNCATE iu9networkslabs.iu9Talgat")
		default:

		}

	}
}
