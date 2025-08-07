package clock

import "fmt"

type Clock struct {
	m int
}

func New(h, m int) Clock {
	m = (h*60 + m) % 1440
	if m < 0 {
		m += 1440
	}
	return Clock{m}
}

func (c Clock) Add(m int) Clock {
	c.m = (c.m + m) % 1440
	if c.m < 0 {
		c.m += 1440
	}
	return c
}

func (c Clock) Subtract(m int) Clock {
	c.m = (c.m - m) % 1440
	if c.m < 0 {
		c.m += 1440
	}
	return c
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.m/60, c.m%60)
}
