package bar

import "time"

type Configs struct {
	updateMask  int
	refreshRate time.Duration
	barWidth    int
	total       int
}

func DefaultConfigs(total int) Configs {
	const (
		defaultMask  = 1023
		defaultRate  = time.Millisecond * 200
		defaultWidth = 50
	)
	return NewConfigs(defaultMask, defaultWidth, total, defaultRate)
}

func NewConfigs(
	updateMask, barWidth, total int,
	refreshRate time.Duration,
) Configs {
	return Configs{
		updateMask:  updateMask,
		refreshRate: refreshRate,
		barWidth:    barWidth,
		total:       total,
	}
}
