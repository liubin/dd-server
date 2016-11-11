package main

import (
	"dd-server/api"
	"dd-server/sink"
	"fmt"
	"github.com/liubin/goutils"
	"github.com/urfave/cli"
	"log"
	"net/http"
	"os"
	"strconv"
)

var flags = []cli.Flag{
	cli.IntFlag{
		Name:  "port",
		Value: 8080,
		Usage: "specific server listening port(8080 by default)",
	},
	cli.StringFlag{
		Name:  "log",
		Usage: "specific output log file, otherwise output to stdout by default",
	},
	cli.StringFlag{
		Name:  "license-validator",
		Usage: "extensional API to validate license",
	},
	cli.StringFlag{
		Name:  "sink-driver",
		Usage: "Sink driver to save the metrics",
	},
	cli.StringSliceFlag{
		Name:  "sink-opts",
		Value: &cli.StringSlice{},
		Usage: "options for the sink driver",
	},
}

// Default handler, do nothing.
func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`Hello dd server!`))
}

func main() {

	app := cli.NewApp()
	app.Name = "dd-server"
	app.Usage = "An open source datadog server"
	app.Version = "0.01"
	app.Author = "bin liu"

	app.Commands = []cli.Command{
		cli.Command{
			Name:   "server",
			Usage:  "start DD server",
			Flags:  flags,
			Action: cmdStartDaemon,
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		panic(fmt.Errorf("Error when executing command: %v", err))
	}
}

func cmdStartDaemon(c *cli.Context) error {

	log_file := c.String("log")
	// save logs to file if set.
	if log_file != "" {
		f, err := os.OpenFile(log_file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

		if err != nil {
			panic(fmt.Sprintf("error opening file: %s", err.Error()))
		}
		defer f.Close()

		log.SetOutput(f)

	}

	driver := c.String("sink-driver")
	log.Printf("driver is %s\n", driver)

	if driver != "" {
		opts := c.StringSlice("sink-opts")
		log.Printf("sink driver opts is %s\n", opts)
		if optsMap, err := goutils.SliceToMap(opts); err != nil {
			return fmt.Errorf("sink opts invalid: %s", opts)
		} else {
			optsMap["sink-driver"] = c.String("sink-driver")
			if e := sink.InitSinkDriver(optsMap); e != nil {
				return e
			}
		}
	}

	licenseValidator := c.String("license-validator")
	log.Printf("license-validator is %s\n", licenseValidator)
	api.SetLicenseValidator(licenseValidator)

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/api/v1/series/", api.SeriesHandler)
	http.HandleFunc("/intake/", api.IntakeHandler)

	port := c.Int("port")

	err := http.ListenAndServe("0.0.0.0:"+strconv.Itoa(port), nil)
	if err != nil {
		fmt.Printf("I think something here could work, but not this")
	}
	return nil
}
