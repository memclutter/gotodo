package config

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net/url"
	"strconv"
)

type Configuration struct {
	Debug        bool
	LogLevel     logrus.Level
	Secret       string
	UrlBase      string
	MailHost     string
	MailPort     int
	MailUsername string
	MailPassword string
}

func (c *Configuration) SetLogLevel(l string) error {
	lvl, err := logrus.ParseLevel(l)
	if err != nil {
		return fmt.Errorf("set log level error: %v", err)
	}
	c.LogLevel = lvl
	return nil
}

func (c *Configuration) SetDsnMail(s string) (err error) {
	u, err := url.Parse(s)
	if err != nil {
		return fmt.Errorf("set dsn mail error: %v", err)
	}
	c.MailHost = u.Host
	c.MailPort, err = strconv.Atoi(u.Port())
	if err != nil {
		return fmt.Errorf("set dsn mail error: %v", err)
	}
	c.MailUsername = u.User.Username()
	c.MailPassword, _ = u.User.Password()

	return nil
}

var Config Configuration
