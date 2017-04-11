package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"time"

	"github.com/ziutek/rrd"
)

// Exit codes are int values that represent an exit code for a particular error.
const (
	ExitCodeOK    int = 0
	ExitCodeError int = 1 + iota
)

// CLI is the command line object
type CLI struct {
	// outStream and errStream are the stdout and stderr
	// to write message from the CLI.
	outStream, errStream io.Writer
}

type Result struct {
	Name string  `json:"name"`
	Max  float64 `json:"max"`
	Min  float64 `json:"min"`
	Avg  float64 `json:"avg"`
}

type Results []Result

// Run invokes the CLI with the given arguments.
func (cli *CLI) Run(args []string) int {
	var (
		pastMin int

		version bool
	)

	// Define option flag parse
	flags := flag.NewFlagSet(Name, flag.ContinueOnError)
	flags.SetOutput(cli.errStream)

	flags.IntVar(&pastMin, "past_min", 86400, "")
	flags.IntVar(&pastMin, "p", 86400, "(Short)")

	flags.BoolVar(&version, "version", false, "Print version information and quit.")

	// Parse commandline flag
	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeError
	}

	// Show version
	if version {
		fmt.Fprintf(cli.errStream, "%s version %s\n", Name, Version)
		return ExitCodeOK
	}
	results := Results{}
	for _, file := range flags.Args() {
		inf, err := rrd.Info(file)
		if err != nil {
			log.Fatal(err)
		}

		end := time.Unix(int64(inf["last_update"].(uint)), 0)
		start := end.Add(time.Duration(pastMin*-1) * time.Minute)

		r, err := rrd.Fetch(file, "MAX", start, end, 1*time.Second)
		if err != nil {
			log.Fatal(err)
		}

		var max, min float64
		var avgv float64
		var avgc float64
		for i := 0; i < len(r.DsNames); i++ {
			for n := 0; n < r.RowCnt; n++ {
				v := math.Floor(r.ValueAt(i, n))

				if v > 0 {
					avgv += v
					if max < v {
						max = v
					}

					if min == 0 || v < min {
						min = v
					}
					avgc++
				}
			}
		}
		if max > 0 {
			results = append(results, Result{
				Name: r.Filename,
				Max:  max,
				Min:  min,
				Avg:  math.Floor(avgv / avgc),
			})
		}

	}
	j, err := json.Marshal(results)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(j))
	return ExitCodeOK
}
