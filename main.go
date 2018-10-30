package main

import (
	"wkcMonitor"
	"strconv"
	"fmt"
	"flag"
)

func main() {
	req := &wkcMonitor.WebContext{}
	var sp string
	var bp string
	flag.BoolVar(&req.IsMonitorBuyPrice, "b", false, "")
	flag.BoolVar(&req.IsMonitorSellPrice, "s", false, "")
	flag.StringVar(&req.MysqlPwd, "mysql", "", "")
	flag.StringVar(&req.EmailCode, "mail", "", "")
	flag.StringVar(&sp, "sp", "", "")
	flag.StringVar(&bp, "bp", "", "")
	flag.Parse()
	var err error
	req.ExpectSellPrice, err = strconv.ParseFloat(sp, 64)
	if err != nil || req.ExpectSellPrice == 0 {
		fmt.Println("swvt9c5bu3 wrong expect sell price")
		return
	}
	req.ExpectBuyPrice, err = strconv.ParseFloat(bp, 64)
	if err != nil || req.ExpectBuyPrice == 0 {
		fmt.Println("mmf5enc5p6 wrong expect buy price")
		return
	}
	wkcMonitor.WebRun(req)
}
