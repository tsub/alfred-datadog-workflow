package main

import (
	"flag"

	"github.com/deanishe/awgo"
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

	wf.NewItem("First result!!").Subtitle("subtitle")

	wf.WarnEmpty("No matching", "Try a different query")
	wf.SendFeedback()
}

func main() {
	aw.Run(run)
}
