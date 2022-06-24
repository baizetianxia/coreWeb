package main

import (
	"time"

	"github.com/baizetianxia/coreWeb/framework"
)

// func UserLoginController(c *framework.Context) error {
// 	c.SetStatus(200).Json("ok, UserLoginController")
// 	return nil
// }

func UserLoginController(c *framework.Context) error {
	foo, _ := c.QueryString("foo", "def")
	// 等待10s才结束执行
	time.Sleep(10 * time.Second)
	// 输出结果
	c.SetOkStatus().Json("ok, UserLoginController: " + foo)
	return nil
}
