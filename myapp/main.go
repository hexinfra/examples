// This is an example showing how to use the Hemi engine to develop applications.

package main

import (
	"github.com/hexinfra/gorox/hemi/contrib/routers/simple"
	"github.com/hexinfra/gorox/hemi/procman"

	. "github.com/hexinfra/gorox/hemi"
)

const usage = `
MyApp (%s)
================================================================================

  myapp [ACTION] [OPTIONS]

ACTION
------

  serve      # start as server
  check      # dry run to check config
  help       # show this message
  version    # show version info
  advise     # show how to optimize current platform
  pids       # call server to report pids of leader and worker
  stop       # tell server to exit immediately
  quit       # tell server to exit gracefully
  leader     # call leader to report its info
  recmd      # tell leader to reopen its cmdui interface
  reweb      # tell leader to reopen its webui interface
  rework     # tell leader to restart worker gracefully
  worker     # call worker to report its info
  reload     # tell worker to reload config
  cpu        # tell worker to perform cpu profiling
  heap       # tell worker to perform heap profiling
  thread     # tell worker to perform thread profiling
  goroutine  # tell worker to perform goroutine profiling
  block      # tell worker to perform block profiling

  Only one action is allowed at a time.
  If ACTION is not specified, the default action is "serve".

OPTIONS
-------

  -debug  <level>   # debug level (default: 0, means disable. max: 2)
  -target <addr>    # leader address to tell or call (default: 127.0.0.1:9527)
  -cmdui  <addr>    # listen address of leader cmdui (default: 127.0.0.1:9527)
  -webui  <addr>    # listen address of leader webui (default: 127.0.0.1:9528)
  -myrox  <addr>    # myrox to use. "-cmdui" and "-webui" will be ignored if set
  -config <config>  # path or url to worker config file
  -single           # run server in single mode. only a process is started
  -daemon           # run server as daemon (default: false)
  -base   <path>    # base directory of the program
  -logs   <path>    # logs directory to use
  -temp   <path>    # temp directory to use
  -vars   <path>    # vars directory to use
  -out    <path>    # daemon's stdout file (default: myapp.log in logs dir)
  -err    <path>    # daemon's stderr file (default: myapp.err in logs dir)

  "-debug" applies to all actions.
  "-target" applies to telling and calling actions only.
  "-cmdui" applies to "serve" and "recmd".
  "-webui" applies to "serve" and "reweb".
  Other options apply to "serve" only.

`

func main() {
	procman.Main("myapp", usage, 0, "127.0.0.1:9527", "127.0.0.1:9528")
}

func init() {
	RegisterHandlet("myHandlet", func(name string, stage *Stage, app *App) Handlet {
		h := new(myHandlet)
		h.onCreate(name, stage, app)
		return h
	})
}

// myHandlet
type myHandlet struct {
	Handlet_
	stage *Stage
	app   *App
}

func (h *myHandlet) onCreate(name string, stage *Stage, app *App) {
	h.MakeComp(name)
	h.stage = stage
	h.app = app

	r := simple.New()

	r.Map("/foo", h.handleFoo)

	h.UseRouter(h, r)
}
func (h *myHandlet) OnShutdown() {
	h.app.SubDone()
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
