package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joerx/gojournal/pkg/handlers"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "address",
			Value: ":8000",
			Usage: "Address to bind to",
		},
	}

	app.Action = startServer

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func startServer(c *cli.Context) error {
	address := c.String("address")

	http.HandleFunc("/", handlers.Index)
	http.HandleFunc("/entries", handlers.GetEntries)

	log.Printf("Listening on %s", address)

	return http.ListenAndServe(address, nil)
}
