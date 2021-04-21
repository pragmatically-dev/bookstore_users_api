package dateutils

import "time"

const (
	apiDateLayout = "2006-01-02 15:04"
)

func GetNow() time.Time  {
	return time.Now().UTC()
}


func GetNowString() string {
	now := GetNow().Format(apiDateLayout)
	return now
}