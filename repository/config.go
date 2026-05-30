package repository

import (
	"os"
	"time"
)

type config struct {
	Addr         string
	Password     string
	PasswordFile string
	DB           int
	Timeout      time.Duration
}

func (c *config) initialize() {
	if c.Timeout == 0 {
		c.Timeout = time.Second * 15
	}
	if c.Password == "" && c.PasswordFile != "" {
		password, err := os.ReadFile(c.PasswordFile)
		if err != nil {
			panic(err)
		}
		c.Password = string(password)
	}
}
