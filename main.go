package main

import (
	"github.com/turbot/steampipe-plugin-mongodbatlas/mongodbatlas"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		PluginFunc: mongodbatlas.Plugin})
}
