/*
@author :wade
EMAIL:wangnmhb@live.com
*/
package main

import (
    "github.com/godlet-io/HowGodletWorks/exp01/server"
)

func main() {
    var httpServer server.HttpServer;
    httpServer.Await();
}
