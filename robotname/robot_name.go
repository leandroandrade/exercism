package robotname

import (
	"math/rand"
	"fmt"
	"time"
)

type Robot string

var nameRobot = make(map[*Robot]bool)

func init() {
	rand.Seed(time.Now().Unix())
}

func (r *Robot) Name() string {
	if nameRobot[r] {
		return string(*r)
	}
	*r = Robot(string(randomValue('A', 'Z')) + string(randomValue('A', 'Z')) +
		fmt.Sprintf("%03d", randomValue(100, 999)))

	nameRobot[r] = true
	return string(*r)
}
func (r *Robot) Reset() {
	nameRobot[r] = false
}

func randomValue(min, max int) int {
	return rand.Intn(max-min) + min
}
