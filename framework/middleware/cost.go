package middleware

import (
	"log"
	"time"

	"github.com/baizetianxia/coreWeb/framework"
)

func Cost() framework.ControllerHandler {
	return func(c *framework.Context) error {
		start := time.Now()

		c.Next()

		end := time.Now()

		cost := end.Sub(start)

		log.Panicf("api uri: %v,cost: %v", c.GetRequest().RequestURI, cost.Seconds())

		return nil
	}
}
