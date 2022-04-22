package main

import (
	"groupcache-test/http_pool"
	"groupcache-test/local_utils"
	"net/http"
)

func main() {
	gRouter := http_pool.NewHttpPool(local_utils.GCFG.Server.Id, "groupcache-test")
	local_utils.Logf(local_utils.GCFG.Server.Id, local_utils.GCFG.Server.Addr, " server  starting")
	srvErr := http.ListenAndServe(local_utils.GCFG.Server.Addr, gRouter)
	if srvErr != nil {
		local_utils.Logf(srvErr.Error())
	}
	local_utils.Logf(local_utils.GCFG.Server.Id, local_utils.GCFG.Server.Addr, " server shutdown")
}
