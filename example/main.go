package main

import (
	"time"

	"github.com/releaseband/go-progress-bar/bar"
)

func main() {
	const total = 10000
	pb := bar.NewProgressBar(bar.DefaultConfigs(total))

	for i := 0; i < total; i++ {
		pb.Update(i)
		time.Sleep(1 * time.Millisecond)
	}

	pb.Finish()
}
