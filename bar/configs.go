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
		defaultMask = 1023
		defaultRate = time.Millisecond * 200
	)
	return NewConfigs(defaultMask, total, defaultRate)
}

func NewConfigs(
	updateMask, total int,
	refreshRate time.Duration,
) Configs {
	return Configs{
		updateMask:  updateMask,
		refreshRate: refreshRate,
		total:       total,
	}
}
