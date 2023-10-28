package main

import (
	"fmt"

	"github.com/choigonyok/goopt/pkg/manifest"
	"github.com/choigonyok/goopt/pkg/middleware"
	"github.com/choigonyok/goopt/pkg/parser"
	"github.com/choigonyok/goopt/pkg/server"
	"github.com/choigonyok/goopt/pkg/watcher"
)

func main() {
	m := manifest.Manifest{Extension: "yml", Name: "test-ap", Path: "../../../.acvctl"}
	w, _ := watcher.New()
	w.WatchFile(m)
	defer w.Close()

	p := parser.New(w)
	go func() {
		for {
			event := <-w.GetEventChannel()
			if event.Op.String() == "WRITE" {
				p.Parse(event)
				fmt.Println(p.GetManifestFromPool(0))
			}
		}
	}()

	middlewares := middleware.New()
	middlewares.AllowOrigin("*")
	middlewares.AllowMethod("GET", "POST", "DELETE", "PUT")
	middlewares.AllowHeader("Origin", "X-Requested-With", "Content-Type", "Accept")
	middlewares.AllowCredential()

	s := server.New(middlewares)
	s.Start()
}
