package main

import (
	"github.com/releaseband/go-progress-bar/bar"
	"time"
)

func main() {
	const total = 10_000
	pb := bar.NewProgressBar(bar.DefaultConfigs(total))

	for i := 0; i < total; i++ {
		pb.Update(i)
		time.Sleep(1 * time.Millisecond)
	}

	pb.Finish()
}
