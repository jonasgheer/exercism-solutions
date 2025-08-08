package clock

import "fmt"

type Clock struct {
	m int
}

const dayInMinutes = 1440

func New(h, m int) Clock {
	m = (h*60 + m) % dayInMinutes
	if m < 0 {
		m += dayInMinutes
	}
	return Clock{m}
}

func (c Clock) Add(m int) Clock {
	c.m = (c.m + m) % dayInMinutes
	if c.m < 0 {
		c.m += dayInMinutes
	}
	return c
}

func (c Clock) Subtract(m int) Clock {
	c.m = (c.m - m) % dayInMinutes
	if c.m < 0 {
		c.m += dayInMinutes
	}
	return c
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.m/60, c.m%60)
}
