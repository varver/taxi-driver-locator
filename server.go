package main

import (
	"runtime"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"gitlab.com/varver/wmd/config"
	"gitlab.com/varver/wmd/controllers"
	"gitlab.com/varver/wmd/db"
)

func main() {
	// parallel process to run
	runtime.GOMAXPROCS(config.Setting.GoMaxProcs)
	m := martini.Classic()
	m.Map(db.Database())
	m.Use(render.Renderer())

	dc := controllers.NewDriverController(db.Database())

	m.Get("/drivers", dc.FetchDrivers)
	m.Put("/driver/(?P<id>[0-9]+)/location", dc.SaveDriverLocation)
	m.Run()

}
