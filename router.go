package main

import "github.com/baizetianxia/coreWeb/framework"

func registerRouter(core *framework.Core) {
	core.Get("foo", FooControllerHandler)
}
