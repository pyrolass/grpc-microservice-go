package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func LogRequest(c *fiber.Ctx) error {

	method := c.Method()
	uri := c.OriginalURL()
	ip := c.IP()

	go func(start time.Time) {

		logrus.WithFields(logrus.Fields{
			"method": method,
			"uri":    uri,
			"ip":     ip,
			"took":   time.Since(start).String(),
		}).Info("Request")
	}(time.Now())

	return c.Next()

}
