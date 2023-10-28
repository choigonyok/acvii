package main

import (
	"fmt"

	"github.com/choigonyok/acvii/pkg/manifest"
	"github.com/choigonyok/acvii/pkg/middleware"
	"github.com/choigonyok/acvii/pkg/parser"
	"github.com/choigonyok/acvii/pkg/server"
	"github.com/choigonyok/acvii/pkg/watcher"
)

const (
	authrizationPolicyManifestExtension = "yml"
	authrizationPolicyManifestName      = "test-ap"
	authrizationPolicyManifestPath      = "../../../.acvctl"
)

func main() {
	mamifest := manifest.Manifest{
		Extension: authrizationPolicyManifestExtension,
		Name:      authrizationPolicyManifestName,
		Path:      authrizationPolicyManifestPath,
	}
	watcher, _ := watcher.New()
	watcher.WatchFile(mamifest)
	defer watcher.Close()

	parser := parser.New(watcher)
	go func() {
		for {
			event := <-watcher.GetEventChannel()
			if event.Op.String() == "WRITE" {
				parser.Parse(event)
				fmt.Println(parser.GetManifestFromPool(0))
			}
		}
	}()

	middleware := middleware.New()
	middleware.AllowOrigin("*")
	middleware.AllowMethod("GET", "POST", "DELETE", "PUT")
	middleware.AllowHeader("Origin", "X-Requested-With", "Content-Type", "Accept")
	middleware.AllowCredential()

	s := server.New(middleware)
	s.Start()
}
