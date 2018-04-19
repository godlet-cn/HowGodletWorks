/*
@author :wade
EMAIL:wangnmhb@live.com
*/
package main

import (
	"github.com/godlet-cn/HowGodletWorks/exp01/server"
)

func main() {
	var httpServer server.HttpServer
	httpServer.Await()
}
