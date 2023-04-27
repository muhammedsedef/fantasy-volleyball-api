package utils

import "time"

func EpochNow() (int64, error) {
	return time.Now().UnixNano() / int64(time.Millisecond), nil
}
