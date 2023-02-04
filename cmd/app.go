package cmd

import (
	"shortener/internals/adapters/framework/driven/db"
	"shortener/internals/adapters/framework/driving/web"
)

func Start() {
	dbAdapter := db.DbAdapter()
	dbAdapter.Connect()
	dbAdapter.Migration()

	webAdapter := web.WebAdapter()
	webAdapter.Run()
}
