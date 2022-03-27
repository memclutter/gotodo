package config

import (
	"fmt"
	"net/url"
	"strconv"
)

type Configuration struct {
	Secret       string
	UrlBase      string
	MailHost     string
	MailPort     int
	MailUsername string
	MailPassword string
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
