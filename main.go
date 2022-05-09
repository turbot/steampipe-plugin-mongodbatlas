package main

import (
	"log"

	"github.com/turbot/steampipe-plugin-mongodbatlas/mongodbatlas"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
)

func main() {

	defer func() {
		r := recover()
		if r != nil {
			log.Println("[ERROR] ", r)
		}
	}()

	plugin.Serve(&plugin.ServeOpts{
		PluginFunc: mongodbatlas.Plugin})
}
