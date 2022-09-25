package wait

import (
	"fmt"
	"testing"
	"time"
)

func Test_Basic(t *testing.T) {
	On(t, For1s, func(s *Status) {
		s.Interval = 100 * time.Millisecond
		// do stuff
		s.Complete = false
		fmt.Println("iteration", s.Iteration())
	})

}
