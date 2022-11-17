package g

import "time"

const (
	TimeFormat = "2006-01-02 15:04:05"
)
const (
	TimeExpireToken  = time.Minute * 1
	TimeRefreshToken = 1 // uint same as TimeExpireToken
)
