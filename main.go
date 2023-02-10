package main

import (
	"log"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/urfave/cli/v2"
)

var app = &cli.App{
	Name:  "serve",
	Usage: "local static file server",
	Flags: []cli.Flag{
		&cli.PathFlag{
			Name:        "dir",
			Usage:       "Serve files from the specified directory",
			Aliases:     []string{"d"},
			Value:       ".",
			Destination: &dir,
		},
		&cli.IntFlag{
			Name:        "port",
			Usage:       "Bind server to custom port",
			Aliases:     []string{"p"},
			Value:       8888,
			Destination: &port,
		},
		&cli.BoolFlag{
			Name:        "share",
			Usage:       "Use host 0.0.0.0 instead of localhost",
			Aliases:     []string{"s"},
			Value:       false,
			Destination: &share,
		},
	},
	Action: action,
}

var (
	dir   cli.Path
	share bool
	port  int
)

func action(c *cli.Context) error {
	return server(dir).Start(addr(share, port))
}

func server(dir string) *echo.Echo {
	e := echo.New()
	e.Use(
		middleware.Recover(),
		middleware.Logger(),
	)
	e.Static("/", dir)
	return e
}

func addr(share bool, port int) string {
	return host(share) + ":" + strconv.Itoa(port)
}

func host(share bool) string {
	if share {
		return "0.0.0.0"
	}
	return "localhost"
}

func main() {
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
