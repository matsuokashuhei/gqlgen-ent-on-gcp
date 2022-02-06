package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/matsuokashuhei/landin/ent"
)

func main() {
	client, err := ent.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DATABASE")))
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}
	defer client.Close()

	p := flag.Bool("print", false, "printing")
	flag.Parse()

	ctx := context.Background()
	if *p {
		f, err := os.Create("migrate.sql")
		if err != nil {
			log.Fatalf("crreate migrate file: %v", err)
		}
		defer f.Close()

		if err := client.Schema.WriteTo(ctx, f); err != nil {
			log.Fatalf("failed printing schema changes: %v", err)
		}
	} else {
		if err := client.Debug().Schema.Create(ctx); err != nil {
			log.Fatalf("failed creating schema resources: %v", err)
		}
	}

}
