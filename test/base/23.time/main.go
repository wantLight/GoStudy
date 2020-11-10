package main

import "fmt"
import "time"

// 摘自：https://www.yiibai.com/go/golang-time.html
// 摘自：https://gobyexample.com/time
// 摘自：https://www.yiibai.com/go/golang-epoch.html
// 摘自：https://gobyexample.com/epoch

func main() {
	p := fmt.Println

	// We'll start by getting the current time.
	now := time.Now()
	p(now)

	// You can build a `time` struct by providing the
	// year, month, day, etc. Times are always associated
	// with a `Location`, i.e. time zone.
	then := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

	// You can extract the various components of the time
	// value as expected.
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	// The Monday-Sunday `Weekday` is also available.
	p(then.Weekday())

	// These methods compare two times, testing if the
	// first occurs before, after, or at the same time
	// as the second, respectively.
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	// The `Sub` methods returns a `Duration` representing
	// the interval between two times.
	diff := now.Sub(then)
	p(diff)

	// We can compute the length of the duration in
	// various units.
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	// You can use `Add` to advance a time by a given
	// duration, or with a `-` to move backwards by a
	// duration.
	p(then.Add(diff))
	p(then.Add(-diff))

	// Use `time.Now` with `Unix` or `UnixNano` to get
	// elapsed time since the Unix epoch in seconds or
	// nanoseconds, respectively.
	secs := now.Unix()
	nanos := now.UnixNano()
	fmt.Println(now)

	// Note that there is no `UnixMillis`, so to get the
	// milliseconds since epoch you'll need to manually
	// divide from nanoseconds.
	millis := nanos / 1000000
	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nanos)

	// You can also convert integer seconds or nanoseconds
	// since the epoch into the corresponding `time`.
	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))
}
