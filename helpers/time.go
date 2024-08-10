package helpers

import "time"

type Time struct{}

type TimeInterface interface {
	NowTime() time.Time
}

func InitTime() TimeInterface {
	return &Time{}
}

func (t *Time) NowTime() time.Time {
	return time.Now()
}
