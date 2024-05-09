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
		Name:    "ahab",
		Version: version,
		Usage:   "tail the whale!",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "since",
				Value: "10m",
				Usage: "Show logs since timestamp [e.g. \"2013-01-02T13:23:37Z\"] or relative [e.g. \"42m\" for 42 minutes]",
			},
		},
		Action: func(cCtx *cli.Context) error {
			containers := cCtx.Args()
			numWorkers := containers.Len()
			if numWorkers == 0 {
				return cli.Exit("you must pass one or more containers to tail!", 2)
			}
			tailer, err := utils.NewLocalDockerTailer(cCtx.String("since"))
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

			// wait for all workers to finish
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
