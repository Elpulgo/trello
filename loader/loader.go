package loader

import (
	"time"

	"github.com/briandowns/spinner"
)

var (
	commandLoader *spinner.Spinner
)

func Run() {
	commandLoader = spinner.New(spinner.CharSets[36], 50*time.Millisecond)
	commandLoader.Start()
	time.Sleep(1 * time.Second)
}

func End() {
	commandLoader.Stop()
}
