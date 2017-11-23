package main

import (
	"github.com/deanishe/awgo"
)

func run() {
	wf := aw.New()

	wf.NewItem("First result!!").Subtitle("subtitle")

	wf.WarnEmpty("No matching", "Try a different query")
	wf.SendFeedback()
}

func main() {
	aw.Run(run)
}
