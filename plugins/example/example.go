package example

import (
	"github.com/chenhg5/go-admin/context"
	"github.com/chenhg5/go-admin/plugins"
	"github.com/chenhg5/go-admin/modules/config"
)

type Example struct {
	app *context.App
}

var Plug = new(Example)

var Config config.Config

func (example *Example) InitPlugin() {
	cfg := config.Get()

	Config = cfg
	Config.PREFIX = "/" + Config.PREFIX

	Plug.app = InitRouter(Config.PREFIX)

}

func NewExample() *Example {
	return Plug
}

func (example *Example) GetRequest() []context.Path {
	return example.app.Requests
}

func (example *Example) GetHandler(url, method string) context.Handler {
	return plugins.GetHandler(url, method, &example.app.HandlerList)
}
