package main

import (
	"os"
  "fmt"

  "database/sql"
	"github.com/codegangsta/cli"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	app := cli.NewApp()
	app.Name = "gchist"
	// app.Version = Version
	app.Usage = ""
	app.Author = "Jun Uchino"
	app.Email = "system.jun@gmail.com"
	app.Action = doMain
	app.Run(os.Args)
}

func doMain(c *cli.Context) {

  db, err := sql.Open("sqlite3", "/Users/juchino/Library/Application Support/Google/Chrome/Default/History")
  if err != nil {
    panic(err)
  }

  rows, err := db.Query("SELECT id, title FROM urls")
  if err != nil {
    panic(err)
  }

  for rows.Next(){
    var id    string
    var title string

    err = rows.Scan(&id, &title)
    fmt.Println(id + " " + title)
  }

  db.Close()
}
