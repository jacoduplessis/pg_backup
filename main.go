package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
	"os/exec"
	"time"
)

func main() {
	dbURL := os.Getenv("DATABASE_URL")

	if dbURL == "" {

		dbURL = "postgres://postgres:postgres@127.0.0.1/postgres?sslmode=disable"
		fmt.Println("[info] DATABASE_URL environment variable not provided, using default")
	}

	db, err := sql.Open("postgres", dbURL)

	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query(`SELECT datname FROM pg_database WHERE datistemplate = false;`)

	if err != nil {
		log.Fatal(err)
	}

	var names []string

	for rows.Next() {

		var name string

		if err := rows.Scan(&name); err != nil {
			log.Fatal(err)
		}

		names = append(names, name)

	}

	dt := time.Now()
	timeString := dt.Format("2006-02-01")

	for _, name := range names {

		outFile := fmt.Sprintf("%s_%s.pg_dump", name, timeString)

		cmd := exec.Command("pg_dump", "-F", "c", "-d", dbURL, "-f", outFile)

		out, err := cmd.Output()

		if err != nil {
			log.Fatal(err)
		}

		_, _ = os.Stdout.Write(out)
	}

}
