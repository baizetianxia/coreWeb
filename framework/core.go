package framework

import (
	"log"
	"net/http"
)

type Core struct {
	router map[string]ControllerHandler
}

func NewCore() *Core {
	return &Core{router: map[string]ControllerHandler{}}
}

func (c *Core) Get(url string, handle ControllerHandler) {
	c.router[url] = handle
}

func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	log.Panic("core.serveHTTP")
	ctx := NewContext(request, response)

	//router selector
	router := c.router["foo"]
	if router == nil {
		return
	}

	log.Panicln("core.router")

	router(ctx)
}
