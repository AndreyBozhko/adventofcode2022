package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

var (
	year int
	day  int
	part string
)

var (
	errEmptyAnswer    = errors.New("answer cannot be empty")
	errSessionMissing = errors.New("session cookie not set")
)

var app = &cli.App{
	Usage: "Advent of Code CLI",
	Commands: []*cli.Command{
		getCliCmd,
		submitCliCmd,
	},
}

var getCliCmd = &cli.Command{
	Name:  "get",
	Usage: "Gets problem input for a specific day",
	Flags: []cli.Flag{yearFlag, dayFlag, sessionFlag},
	Action: func(ctx *cli.Context) error {
		return GetInput(&InputRequest{
			year, day, fmt.Sprintf("inputs/%d/input%02d.txt", year, day),
		})
	},
}

var submitCliCmd = &cli.Command{
	Name:      "submit",
	Usage:     "Submits answer for a specific problem",
	ArgsUsage: "answer",
	Flags:     []cli.Flag{yearFlag, dayFlag, partFlag, sessionFlag},
	Action: func(ctx *cli.Context) error {
		ans := ctx.Args().First()
		if ans == "" {
			return errEmptyAnswer
		}

		return SubmitAnswer(
			ans,
			&SubmissionRequest{
				year,
				day,
				part,
			})
	},
}

var yearFlag = &cli.IntFlag{
	Name:        "year",
	Aliases:     []string{"y"},
	Destination: &year,
	Value:       2022,
	Required:    false,
	Action: func(ctx *cli.Context, i int) error {
		if i < 2019 || i > 2023 {
			return fmt.Errorf("invalid year: %d, must be between 2019 and 2023", i)
		}
		return nil
	},
}

var dayFlag = &cli.IntFlag{
	Name:        "day",
	Aliases:     []string{"d"},
	Destination: &day,
	DefaultText: "n/a",
	Required:    true,
	Action: func(ctx *cli.Context, i int) error {
		if i < 1 || i > 25 {
			return fmt.Errorf("invalid day: %d, must be between 1 and 25", i)
		}
		return nil
	},
}

var partFlag = &cli.StringFlag{
	Name:        "part",
	Aliases:     []string{"p"},
	Destination: &part,
	Required:    true,
	Action: func(ctx *cli.Context, s string) error {
		if s == "A" || s == "B" {
			return nil
		}
		return fmt.Errorf("invalid problem part: %s, must be A or B", s)
	},
}

var sessionFlag = &cli.StringFlag{
	Name:        "session",
	Aliases:     []string{"s"},
	Destination: &Session,
	Required:    true,
	EnvVars:     []string{"SESSION"},
	FilePath:    ".session",
	Action: func(ctx *cli.Context, s string) error {
		if s == "" {
			return errSessionMissing
		}
		return nil
	},
}

func main() {
	log.SetFlags(0)
	log.SetPrefix(time.Now().Format("2006-01-02 15:04:05 MST") + ": ")

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
