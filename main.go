package main

import (
	"flag"
	"fmt"

	"github.com/deanishe/awgo"
	"github.com/zorkian/go-datadog-api"
)

func run() {
	wf := aw.New()

	apiKey := flag.String("apikey", "", "Datadog api key")
	appKey := flag.String("appkey", "", "Datadog application key")
	flag.Parse()

	if *apiKey == "" || *appKey == "" {
		wf.NewWarningItem("Datadog api key or application key is not set!", "subtitle")
		wf.SendFeedback()
		return
	}

	d := datadog.NewClient(*apiKey, *appKey)

	switch flag.Arg(0) {
	case "dashboard":
		dashboards, err := d.GetDashboards()
		if err != nil {
			wf.FatalError(err)
		}

		for _, dash := range dashboards {
			url := fmt.Sprintf("https://app.datadoghq.com/dash/%d/datadog", dash.GetId())
			wf.NewItem(dash.GetTitle()).Subtitle(url).Arg(url).Valid(true)
		}
	}

	wf.WarnEmpty("No matching", "Try a different query")
	wf.SendFeedback()
}

func main() {
	aw.Run(run)
}
