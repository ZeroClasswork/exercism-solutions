// package gigasecond implements a function that adds 
// 	  a gigasecond to a given moment in time.
package gigasecond

// import path for the time package from the standard library
import "time"

// AddGigasecond adds a gigsecond of time to a given
//    moment in time.
func AddGigasecond(t time.Time) time.Time {
	t = t.Add(time.Duration(1000000000*1000000000))
	return t
}
