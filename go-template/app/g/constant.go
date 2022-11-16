package g

import "time"

const (
	TimeFormat = "2006-01-02 15:04:05"
)
const (
	TimeExpireToken  = time.Minute * 7
	TimeRefreshToken = 5 // uint same as TimeExpireToken
)
