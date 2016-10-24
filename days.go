package days

import (
	"time"
)

// Returns date with the time portion of the day truncated
func Trunc(tm time.Time) (time.Time) {
	_, offset := tm.Zone()
	seconds := tm.Unix()
	return time.Unix(seconds-(seconds+int64(offset))%86400, 0)
}

// Returns today's time to midnight
func Today() (time.Time) {
	return Trunc(time.Now())
}

// Returns tomorrow's time by adding one day to the given time
func Tomorrow(tm time.Time) (time.Time) {
	return Add(tm, 1)
}

// Returns tomorrow's time by subtracting one day from the given time
func Yesterday(tm time.Time) (time.Time) {
	return Sub(tm, 1)
}

func Days(tm time.Time) int {
	_, offset := tm.Zone()
	return int((tm.Unix() + int64(offset)) / int64(86400))
}

// Calculates the difference between two times in days
func Between(t1, t2 time.Time) (int) {
	return Days(t1) - Days(t2)
}

// Adds "n" days to a given time
func Add(tm time.Time, n int) (ret time.Time) {
	ret = Trunc(tm)
	if n != 0 {
		ret = time.Unix(ret.Unix() + int64(86400*n), 0)
	}
	return
}

// Subtracts "n" days to a given time
func Sub(tm time.Time, n int) (ret time.Time) {
	ret = Trunc(tm)
	if n != 0 {
		ret = time.Unix(ret.Unix() - int64(86400 * n), 0)
	}
	return
}
