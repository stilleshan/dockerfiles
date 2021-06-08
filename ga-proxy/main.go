package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/giuem/ga-proxy/server"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "ga-proxy"
	app.HideVersion = true
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "ip, i",
			Value:  "127.0.0.1",
			Usage:  "`IP` to listen",
			EnvVar: "IP",
		},
		cli.StringFlag{
			Name:   "port, p",
			Value:  "9080",
			Usage:  "`port` to listen",
			EnvVar: "PORT",
		},
	}

	app.Action = func(c *cli.Context) error {
		server.Run(c.String("ip"), c.String("port"))
		return nil
	}

	app.Commands = []cli.Command{
		cli.Command{
			Name: "ping",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:   "ip, i",
					Value:  "127.0.0.1",
					Usage:  "server `IP`",
					EnvVar: "IP",
				},
				cli.StringFlag{
					Name:   "port, p",
					Value:  "9080",
					Usage:  "server `port`",
					EnvVar: "PORT",
				},
			},
			Action: func(c *cli.Context) error {
				resp, err := http.Get(fmt.Sprintf("http://%v:%v/ping", c.String("ip"), c.String("port")))
				if err != nil {
					return cli.NewExitError(err, 1)
				}
				defer resp.Body.Close()
				if resp.StatusCode != 200 {
					return cli.NewExitError(fmt.Errorf("server returns non-200 status code"), 1)
				}
				return nil
			},
		},
	}

	app.Run(os.Args)
}
