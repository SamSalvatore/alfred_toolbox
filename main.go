package main

// Package is called aw
import (
	"alfred_toolbox/service/time_tool"
	"github.com/deanishe/awgo"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

// Workflow is the main API
var wf *aw.Workflow

func init() {
	// Create a new Workflow using default settings.
	// Critical settings are provided by Alfred via environment variables,
	// so this *will* die in flames if not run in an Alfred-like environment.
	wf = aw.New()
}

// Your workflow starts here
func run() {
	//// Add a "Script Filter" result
	//wf.NewItem("First result!")
	//// Send results to Alfred
	//wf.SendFeedback()
}

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "ts",
				Usage: "时间戳转换工具",
				Action: func(c *cli.Context) error {
					time_tool.ConvertTime(wf, c)
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
