package myapp

import (
	"github.com/hexinfra/gorox/hemi/classic/mappers/simple"

	. "github.com/hexinfra/gorox/hemi"
)

func init() {
	RegisterHandlet("myHandlet", func(name string, stage *Stage, webapp *Webapp) Handlet {
		h := new(myHandlet)
		h.onCreate(name, stage, webapp)
		return h
	})
}

type myHandlet struct {
	Handlet_
}

func (h *myHandlet) onCreate(name string, stage *Stage, webapp *Webapp) {
	h.Handlet_.OnCreate(name, stage, webapp)

	r := simple.New()
	r.Map("/foo", h.handleFoo)
	h.UseMapper(h, r)
}
func (h *myHandlet) OnShutdown() {
	h.Webapp().DecSub()
}

func (h *myHandlet) OnConfigure() {}
func (h *myHandlet) OnPrepare()   {}

func (h *myHandlet) Handle(req Request, resp Response) (next bool) {
	h.Dispatch(req, resp, h.notFound)
	return
}
func (h *myHandlet) notFound(req Request, resp Response) {
	resp.Send("handle not found!")
}

func (h *myHandlet) handleFoo(req Request, resp Response) { // METHOD /foo
	resp.Echo(req.H("user-agent"))
}

func (h *myHandlet) GET_(req Request, resp Response) { // GET /
	resp.Echo("hello, world! ")
	resp.Echo("this is an example application.")
}
func (h *myHandlet) POST_user_login(req Request, resp Response) { // POST /user/login
	resp.Send("what are you doing?")
}
