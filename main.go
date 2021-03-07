package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"egreb.net/todos/database"
	"egreb.net/todos/routes"
)

func main() {
	if err := run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	db := database.Connect()
	if err := database.CreateSchema(db); err != nil {
		log.Fatal("Could not connect to the database", err)
	}
	defer db.Close()

	flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
	var (
		port = flags.Int("port", 8080, "port to listen on")
	)
	if err := flags.Parse(args[1:]); err != nil {
		return err
	}
	addr := fmt.Sprintf("0.0.0.0:%d", *port)
	srv, err := routes.NewServer()
	if err != nil {
		return err
	}
	srv.Routes(db)
	fmt.Printf("server listening on :%d\n", *port)
	return http.ListenAndServe(addr, srv)
}
