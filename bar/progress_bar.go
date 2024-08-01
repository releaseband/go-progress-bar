package bar

import (
	"fmt"
	"strings"
	"time"
)

type Progress struct {
	progressInformation
	Configs
	current int
	speed   float64
}

type progressInformation struct {
	startTime     time.Time
	lastUpdate    time.Time
	nextCheckTime time.Time

	elapsedTime   time.Duration
	remainingTime time.Duration
}

func NewProgressBar(cfg Configs) *Progress {
	now := time.Now()

	return &Progress{
		Configs: cfg,
		progressInformation: progressInformation{
			startTime:     now,
			lastUpdate:    now,
			nextCheckTime: now.Add(cfg.refreshRate),

			elapsedTime:   0,
			remainingTime: 0,
		},
		current: 0,
		speed:   0,
	}
}

func (pb *Progress) Update(current int) {
	pb.current = current
	now := time.Now()

	if (current&pb.updateMask) == 0 || current == pb.total {
		if now.After(pb.nextCheckTime) || current == pb.total {
			pb.updateStats(now)
			pb.render()
			pb.lastUpdate = now
			pb.nextCheckTime = now.Add(pb.refreshRate)
		}
	}
}

func (pb *Progress) updateStats(now time.Time) {
	pb.elapsedTime = now.Sub(pb.startTime)
	pb.speed = float64(pb.current) / pb.elapsedTime.Seconds()

	if pb.speed > 0 {
		remainingCount := pb.total - pb.current
		pb.remainingTime = time.Duration(float64(remainingCount) / pb.speed * float64(time.Second))
	}
}

func (pb *Progress) render() {
	ratio := float64(pb.current) / float64(pb.total)
	percent := int(ratio * 100)

	barLength := int(ratio * float64(pb.barWidth))
	bar := strings.Repeat("█", barLength) + strings.Repeat("░", pb.barWidth-barLength)

	fmt.Printf("\r[%s] %3d%% | %d/%d | %.2f it/s | %v elapsed | %v remaining",
		bar, percent, pb.current, pb.total, pb.speed,
		pb.elapsedTime.Round(time.Second), pb.remainingTime.Round(time.Second),
	)
}

func (pb *Progress) Finish() {
	pb.Update(pb.total)
	fmt.Println()
}
