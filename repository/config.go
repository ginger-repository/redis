package repository

import "time"

type config struct {
	Addr     string
	Password string
	DB       int
	Timeout  time.Duration
}

func (c *config) initialize() {
	if c.Timeout == 0 {
		c.Timeout = time.Second * 15
	}
}
