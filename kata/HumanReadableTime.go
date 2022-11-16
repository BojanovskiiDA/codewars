package kata

import (
	"fmt"
)

func HumanReadableTime(seconds int) string {
	hh := seconds / 3600
	mm := (seconds % 3600) / 60
	ss := seconds % 60
	return fmt.Sprintf("%02d:%02d:%02d", hh, mm, ss) 
}