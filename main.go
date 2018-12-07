package main

import (
	"wkcMonitor"
	"flag"
)

func main() {
	req := &wkcMonitor.WebContext{}
	flag.BoolVar(&req.IsMonitorBuyPrice, "b", false, "")
	flag.BoolVar(&req.IsMonitorSellPrice, "s", false, "")
	flag.StringVar(&req.MysqlPwd, "mysql", "", "")
	flag.Parse()
	wkcMonitor.WebRun(req)
}
