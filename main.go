package main

import (
	"log"
	"os"

	"ahab/utils"

	"github.com/urfave/cli/v2"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	app := &cli.App{
		Name:                 "ahab",
		Version:              version,
		ArgsUsage:            "[container IDs or names]",
		Usage:                "tail the whale!",
		EnableBashCompletion: true,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "since",
				Aliases: []string{"s"},
				Value:   "10m",
				Usage:   "show logs since timestamp [e.g. \"2013-01-02T13:23:37Z\"] or relative [e.g. \"42m\" for 42 minutes]",
			},
			&cli.StringFlag{
				Name:    "until",
				Aliases: []string{"u"},
				Value:   "",
				Usage:   "show logs until timestamp [e.g. \"2013-01-02T13:23:37Z\"] or relative [e.g. \"42m\" for 42 minutes]",
			},
			&cli.BoolFlag{
				Name:    "no-follow",
				Usage:   "don't follow after printing the tails of the given containers",
				EnvVars: []string{"AHAB_NO_FOLLOW"},
			},
			&cli.BoolFlag{
				Name:    "no-timestamps",
				Usage:   "don't print timestamps with log entries",
				EnvVars: []string{"AHAB_NO_TIMESTAMPS"},
			},
			// &cli.StringFlag{
			// 	Name:    "custom-colors",
			// 	Aliases: []string{"c"},
			// 	Usage:   "a comma separated list of custom colors to use for container names",
			// 	EnvVars: []string{"AHAB_CUSTOM_COLORS"},
			// },
		},
		Action: func(cCtx *cli.Context) error {
			containers := cCtx.Args()
			numWorkers := containers.Len()
			if numWorkers == 0 {
				return cli.Exit("you must pass one or more container names or IDs to tail!", 2)
			}
			tailer, err := utils.NewLocalDockerTailer(cCtx.String("since"), cCtx.String("until"), !cCtx.Bool("no-follow"), !cCtx.Bool("no-timestamps"))
			if err != nil {
				log.Fatal(err)
			}
			workers := make(chan int, numWorkers)
			for index := range numWorkers {
				go func(c string) {
					err := tailer.Tail(c, utils.NewLogWriter(c))
					if err != nil {
						log.Fatal(err)
					}
					workers <- 0
				}(containers.Get(index))
			}

			for i := 0; i < numWorkers; i++ {
				<-workers
			}
			close(workers)

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
