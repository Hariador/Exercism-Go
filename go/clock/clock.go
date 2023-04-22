package clock

import (
	"fmt"
)

type Clock struct {
	m int
}

func New(h, m int) Clock {
	m = convMin(m + (60 * h))
	c := Clock{m: m}
	return c
}

func (c Clock) Add(m int) Clock {
	c.m = convMin(c.m + m)
	return c

}

func (c Clock) Subtract(m int) Clock {
	c.m = convMin(c.m - m)
	return c
}

func (c Clock) String() string {

	minutes := c.m % 60
	hours := c.m / 60
	return fmt.Sprintf("%02d:%02d", hours%24, minutes)

}

func convMin(m int) int {
	for m > 1440 {
		m = m - 1440
	}
	for m < 0 {
		m = m + 1440
	}
	if m == 1440 {
		m = 0
	}
	return m
}
