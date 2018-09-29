package temperature

import (
	"time"
)

type Message struct {
	Date time.Time
	Text string
	Name string
}
